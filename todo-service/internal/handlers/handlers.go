package handlers

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"

	"todo-service/internal/repo"
	"todo-service/internal/models"
	
	auth "todo/auth-service/pb"
)

type Handler struct {
	AuthClient auth.AuthServiceClient
	Tdb *sqlx.DB
	Trs *repo.TodoRepoSql
}

func extTask(c echo.Context, task *models.Task, userID *string) (map[string]interface{}, error) {
	uid, ok := c.Get("userID").(string)
	if !ok {
		return nil, echo.NewHTTPError(
			http.StatusUnauthorized, "Invalid user id")
	}
	*userID = uid

	if err := c.Bind(task); err != nil {
		return nil, echo.NewHTTPError(
			http.StatusBadRequest,
			"Invalid request data:\n" + err.Error())
	}

	fields := map[string]interface{}{
		"task_id": task.ID,
		"title": task.Title,
		"category_id": task.Category,
		"done": task.Done,
	}
	for _, field := range []string{"task_id", "title", "category_id", "done"} {
		if value := fields[field]; value != nil {
			return map[string]interface{}{field: value}, nil
		}
	}
	return nil, nil
}

func (h *Handler) AddUser(c echo.Context) error {
	ctx := c.Request().Context()
	var user models.User
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest, "Invalid fields")
	}

	hashRes, err := h.AuthClient.HashPd(ctx, &auth.HashReq{Password: user.PdHash})
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"Hash password error\n" + err.Error())
	}
	user.PdHash = hashRes.Hash

	if err = h.Trs.AddUser(user); err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"Error adding a user to the db\n" + err.Error())
	}

	token, err := h.AuthClient.GenJWT(
		ctx, &auth.JWTReq{UserID: user.Id})
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"Error generating the JWT token\n" + err.Error())
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": token.Token,
	})
}

func (h *Handler) AddOrUpdTask(c echo.Context) error {
	var userID string
	var task models.Task
	if _, err := extTask(c, &task,&userID); err != nil {return err}

	if err := h.Trs.AddOrUpdTask(&task, userID); err != nil {
		return echo.NewHTTPError (
			http.StatusInternalServerError,
			"error adding a task to the db:\n" + err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Task successfully added",
		"task": task,
	})
}

func (h *Handler) GetTask(c echo.Context) error {
	var userID string
	var task models.Task
	found, err := extTask(c, &task, &userID)
	if err != nil {return err}

	tasks, err := h.Trs.GetTask(userID, task, found);
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"error when getting a task from the database:\n"+ err.Error())
	}

	if len(tasks) == 0 {
		return echo.NewHTTPError(
			http.StatusNotFound, "No tasks found")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"tasks": tasks,
	})
}
