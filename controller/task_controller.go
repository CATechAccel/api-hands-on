package controller

import (
	"github.com/labstack/echo"
	"net/http"
)

func ListTasks(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!あああああ")
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
