package main

import (
	"fmt"
	"os"
	"pasar/config"
	"pasar/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	config.ConnectDatabase()
	e := echo.New()
	e = routes.InitRoute(e)
	e.Start(getPort())
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" // Default port if not specified
	}

	return fmt.Sprintf(":%s", port)
}
