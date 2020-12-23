package router

import (
	"github.com/fumist23/api-handson/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Router() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// http://localhost:8080/api/v1/tasks
	tasks := e.Group("/api/v1/tasks")
	tasks.GET("", controller.ListTasks)
	tasks.GET("/:id", controller.FindById)
	tasks.POST("", controller.CreateTask)
	tasks.PUT("/:id", controller.UpdateTask)
	tasks.DELETE("/:id", controller.DeleteTask)

	return e
}
