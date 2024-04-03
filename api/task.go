package api

import (
	"net/http"
	"strconv"
	"togolist/model"

	"github.com/labstack/echo/v4"
)

func TaskHandler(c echo.Context) error {
	tasks, err := model.GetAllTasks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, tasks)
}

func CreateTaskHandler(c echo.Context) error {
	task := new(model.TaskCreation)
	if err := c.Bind(task); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := model.CreateTask(task); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, task)
}

func GetTaskHandler(c echo.Context) error {
	task := new(model.Task)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	task.ID = id
	if err := model.GetTask(task); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, task)
}
