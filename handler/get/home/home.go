package home

import (
	"net/http"

	. "test/db/getall"

	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {

	users := Getall()

	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"users": users,
	})
}
