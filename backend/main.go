package main

import (
	"fmt"
	"net/http"
	"runtime"

	"binalyze-test/configs"
	"binalyze-test/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// _err := godotenv.Load(".env")

	// if _err != nil {
	// 	log.Print(_err)
	// }

	fmt.Println("v", runtime.GOOS)

	configs.ConnectDb()

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome !")
	})

	group := e.Group("/api/v1")

	routes.Routes(group)

	RunSchedule()

	port := fmt.Sprintf(":%s", "1323")

	// Start server
	e.Logger.Fatal(e.Start(port))

	select {}
}
