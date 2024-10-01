package Controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/lucasfmarconi/switcher-api/Cache/CacheServices"
)

type Switch struct {
    Key       string        `json:"key"`
    Value     bool        	`json:"value"`
    ExpiresIn time.Duration `json:"expires_in"` // Duration from now
}

// UpdateSwitchHandler handles the update switch request
func UpdateSwitchHandler(c echo.Context) error {
    var sw Switch
    if err := c.Bind(&sw); err != nil {
        return c.String(http.StatusBadRequest, "Invalid request payload")
    }

    CacheServices.SetCache(sw.Key, fmt.Sprintf("%t", sw.Value))

    // Construct the URL for the created resource
    resourceURL := fmt.Sprintf("%s%s/%s", c.Request().Host, c.Request().RequestURI, sw.Key)

    // Return the key and the URL of the created resource
    return c.JSON(http.StatusAccepted, map[string]string{
        "key": sw.Key,
        "url": resourceURL,
    })
}

// GetSwitchHandler handles the get switch request
func GetSwitchHandler(c echo.Context) error {
	/// TODO: Implement get logic here
    sw := Switch{
        Key:       "exampleKey",
        Value:     true,
        ExpiresIn: 24 * time.Hour,
    }

    return c.JSON(http.StatusOK, sw)
}