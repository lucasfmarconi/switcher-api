package Controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// RootHandler handles the root URL
func RootHandler(c echo.Context) error {
	isTls := c.IsTLS()
	return c.String(http.StatusOK, fmt.Sprintf("Is TLS: %v", isTls))
}