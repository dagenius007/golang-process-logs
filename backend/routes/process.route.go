package routes

import (
	"binalyze-test/handlers"

	"github.com/labstack/echo/v4"
)

func ProcessRoute(e *echo.Group) {
	e.GET("/data", handlers.GetProcess)
	e.GET("/ws", handlers.GetProcessRealTime)
	e.GET("/users", handlers.GetProcessUsers)
	e.GET("/counts", handlers.GetProcessCounts)
	e.GET("/reports", handlers.GetProcessReports)
}
