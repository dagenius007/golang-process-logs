package routes

import (
	"process-logs/handlers"
	"process-logs/setup"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Group, services *setup.ServiceDependencies) {
	handlers.UseProcessRoutes(e.Group("/process"), services)
}
