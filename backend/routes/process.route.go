package routes

import (
	"binalyze-test/handlers"

	"github.com/labstack/echo/v4"
)

func ProcessRoute(e *echo.Group) {
	e.GET("/data", handlers.GetProcess)
	e.GET("/ws", handlers.GetProcessRealTime)
}
