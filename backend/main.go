package main

import (
	"fmt"
	"log"
	"net/http"

	"binalyze-test/routes"
	"binalyze-test/setup"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func main() {
	_err := godotenv.Load(".env")

	if _err != nil {
		log.Print(_err)
	}

	// configs.ConnectDb()

	logger := logrus.New()

	services, err := setup.ConfigureServiceDependencies(logger)
	if err != nil {
		logger.Println("failed to setup service dependencies")
		panic(err)
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome !")
	})

	group := e.Group("/api/v1")

	routes.Routes(group, services)

	go RunSchedule(services)

	port := fmt.Sprintf(":%s", "1323")

	// Start server
	e.Logger.Fatal(e.Start(port))

	select {}
}
