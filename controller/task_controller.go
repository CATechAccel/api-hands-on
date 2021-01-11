package controller

import (
	"github.com/fumist23/api-handson/database"
	"github.com/labstack/echo"
	"golang.org/x/crypto/ssh"
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
	//ここでc.Contextからrequestのbodyのなかのtaskを受け取る
	result, err := database.InsertTask(ctx, task)
}

func UpdateTask(c echo.Context) error {
	return nil
}

func DeleteTask(c echo.Context) error {
	return nil
}
