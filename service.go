package f

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Trigger(c echo.Context) error {
	// business logic
	// ...
	return c.String(http.StatusOK, "Done")
}
