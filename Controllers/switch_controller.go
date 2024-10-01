package Controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/lucasfmarconi/switcher-api/CacheServices"
)

type Switch struct {
    Key       string        `json:"key"`
    Value     bool        	`json:"value"`
    ExpiresIn time.Duration `json:"expires_in"` // Duration from now
}

// UpdateSwitchHandler handles the update switch request
func UpdateSwitchHandler(c echo.Context) error {
    var switchInstance Switch
    if err := c.Bind(&switchInstance); err != nil {
        return c.String(http.StatusBadRequest, "Invalid request payload")
    }

    CacheServices.SetCache(switchInstance.Key, fmt.Sprintf("%t", switchInstance.Value))

    // Construct the URL for the created resource
    resourceURL := fmt.Sprintf("%s%s/%s", c.Request().Host, c.Request().RequestURI, switchInstance.Key)

    // Return the key and the URL of the created resource
    return c.JSON(http.StatusAccepted, map[string]string{
        "key": switchInstance.Key,
        "url": resourceURL,
    })
}

// GetSwitchHandler handles the get switch request
func GetSwitchHandler(c echo.Context) error {
    // Get Switch Struct object from cache
    switchInstance := CacheServices.GetCache(c.Param("key"))
    return c.JSON(http.StatusOK, switchInstance)
}