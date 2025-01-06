package main

import (
	"switcher-api/Controllers"

	"github.com/labstack/echo/v4"
)

func main() {
	// Create a new Echo instance
	echoInstance := echo.New()

	// Define routes
	defineRoutes(echoInstance)

	// Start the server
	echoInstance.Logger.Fatal(echoInstance.Start(":8080"))
}

// defineRoutes sets up the routes for the Echo instance
func defineRoutes(e *echo.Echo) {
	e.GET("/", Controllers.RootHandler)
	e.POST("/switch", Controllers.UpdateSwitchHandler)
	e.GET("/switch/:key", Controllers.GetSwitchHandler)
}