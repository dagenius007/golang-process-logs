package routes

import (
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Group, services *main.ServiceDependencies) {
	ProcessRoute(e)
	handlers.UseProcessRoute(e.Group("/auth"), opts)
}
