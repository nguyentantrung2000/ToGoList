package main

import (
	"net/http"
	"togolist/api"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	v1 := e.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.GET("", api.TasksHandler)
			items.POST("", api.CreateTaskHandler)
			items.GET("/:id", api.GetTaskHandler)
			items.PATCH("/:id", api.UpdateTaskHandler)
			items.DELETE("/:id", api.DeleteTaskHandler)
		}
	}
	e.Start(":8080")
}
