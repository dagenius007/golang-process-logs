package routes

import (
	"binalyze-test/handlers"
	"binalyze-test/setup"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Group, services *setup.ServiceDependencies) {
	handlers.UseProcessRoutes(e.Group("/data"), services)
}
