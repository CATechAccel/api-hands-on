package controller

import (
	"github.com/fumist23/api-handson/database"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

func ListTasks(c echo.Context) error {
	ctx := c.Request().Context()
	tasks, err := database.GetList(ctx)
	if err != nil {
		log.Printf("error: %v", err)
		c.String(http.StatusInternalServerError, "サーバーエラー: tasksの取得")
	}
	return c.JSON(http.StatusOK, tasks)
}

func FindById(c echo.Context) error {
	return nil
}

func CreateTask(c echo.Context) error {
	return nil
}

func UpdateTask(c echo.Context) error {
	return nil
}

func DeleteTask(c echo.Context) error {
	return nil
}
