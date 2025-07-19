package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"todo-service/internal/models"
)

func (h *Handler) AddOrUpdTask(c echo.Context) error {
	var userID string
	var task models.Task
	if _, err := extTask(c, &task, &userID); err != nil {
		return err
	}

	if err := h.Trs.AddOrUpdTask(&task, userID); err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (h *Handler) GetTask(c echo.Context) error {
	var userID string
	var task models.Task

	found, err := extTask(c, &task, &userID)
	if err != nil {
		return err
	}

	tasks, err := h.Trs.GetTask(userID, task, found)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError, err.Error())
	}

	if len(tasks) == 0 {
		return echo.NewHTTPError(
			http.StatusNotFound, "No tasks found")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"tasks": tasks,
	})
}

func (h *Handler) DelTask(c echo.Context) error {
	var userID string
	var task models.Task

	found, err := extTask(c, &task, &userID)
	if err != nil {
		return err
	}

	taskID, ok := found["task_id"].(int)
	if !ok || taskID == 0 {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"Invalid or missing task_id")
	}

	if err := h.Trs.DelTask(userID, taskID); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
