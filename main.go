package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	echoInstance := echo.New()
	echoInstance.GET("/", func(c echo.Context) error {
		isTls := c.IsTLS()

		return c.String(http.StatusOK, fmt.Sprintf("Is TLS: %v", isTls))
	})

	echoInstance.POST("/switch/:id", updateSwitch)
	echoInstance.GET("/switch/:id", getSwitch)

	echoInstance.Logger.Fatal(echoInstance.Start(":8080"))
}

func updateSwitch(c echo.Context) error {
	// Implement your update logic here
	return c.String(http.StatusOK, "Switch updated")

}

func getSwitch(c echo.Context) error {
	// Implement your update logic here
	return c.String(http.StatusOK, fmt.Sprintf("Switch Id: %v", c.Param("id")))

}