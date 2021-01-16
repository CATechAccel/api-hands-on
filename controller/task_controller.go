package controller

import (
	"encoding/json"
	"github.com/fumist23/api-handson/database"
	"github.com/fumist23/api-handson/model"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"strconv"
)

func ListTasks(c echo.Context) error {
	ctx := c.Request().Context()
	tasks, err := database.GetList(ctx)
	if err != nil {
		log.Printf("GetList Error: %v", err)
		return c.String(http.StatusInternalServerError, "サーバーエラー: GetList")
	}
	return c.JSON(http.StatusOK, tasks)
}

func FindById(c echo.Context) error {
	ctx := c.Request().Context()
	id, _ := strconv.Atoi(c.Param("id"))
	task, err := database.GetTask(ctx, id)
	if err != nil {
		log.Printf("GetTask Error: %v", err)
		return c.String(http.StatusInternalServerError, "GetTask Error")
	}
	return c.JSON(http.StatusOK, task)
}

func CreateTask(c echo.Context) error {
	ctx := c.Request().Context()
	var task model.Task
	err := json.NewDecoder(c.Request().Body).Decode(&task)
	_, err = database.InsertTask(ctx, task)
	if err != nil {
		return c.String(http.StatusInternalServerError, "InsertTask Error")
	}
	return c.String(http.StatusCreated, "Created!!")
}

func UpdateTask(c echo.Context) error {
	ctx := c.Request().Context()
	id, _ := strconv.Atoi(c.Param("id"))
	var task model.Task
	err := json.NewDecoder(c.Request().Body).Decode(&task)
	if err != nil {
		log.Printf("failed to decode request body %v", err)
		return c.String(http.StatusInternalServerError, "failed to decode request")
	}
	_, err = database.UpdateTask(ctx, id, task)
	if err != nil {
		return c.String(http.StatusInternalServerError, "UpdateTask Error")
	}
	return c.String(http.StatusCreated, "Updated!!")
}

func DeleteTask(c echo.Context) error {
	ctx := c.Request().Context()
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := database.DeleteTask(ctx, id)
	if err != nil {
		return c.String(http.StatusInternalServerError, "DeleteTask Error")
	}
	return c.String(http.StatusOK, "Deleted!!")
}
