package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	_err := godotenv.Load(".env")

	if _err != nil {
		log.Print(_err)
	}

	// configs.ConnectDb()
	// configs.ConnectRedis()

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Welcome !")
	// })

	// group := e.Group("/api/v1")

	// routes.Routes(group)

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))

	// Start server
	e.Logger.Fatal(e.Start(port))
}
