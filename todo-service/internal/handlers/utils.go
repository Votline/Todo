package handlers

import (
	"reflect"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"

	"todo-service/internal/repo"
	"todo-service/internal/models"

	auth "todo/auth-service/pb"
)

type Handler struct {
	AuthClient auth.AuthServiceClient
	Tdb        *sqlx.DB
	Trs        *repo.TodoRepoSql
}

func getValue(v interface{}) interface{} {
	if v == nil {
		return nil
	}

	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return nil
		}
		return val.Elem().Interface()
	}
	return v
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
			http.StatusBadRequest, err.Error())
	}

	fields := map[string]interface{}{
		"task_id":     getValue(task.ID),
		"title":       getValue(task.Title),
		"category_id": getValue(task.Category),
		"done":        getValue(task.Done),
	}
	for _, field := range []string{"task_id", "title", "category_id", "done"} {
		if value := fields[field]; value != nil {
			return map[string]interface{}{field: value}, nil
		}
	}
	return nil, echo.NewHTTPError(
		http.StatusBadRequest, "Invalid or missing fields")
}
