package form

import (
	"fmt"
	"net/http"

	. "test/db/saveuser"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name  string `json:"name" form:"name" validate:"required"`
	Email string `json:"email" form:"email" validate:"required,email"`
}

func Form(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		fmt.Println(err)
	}

	Saveuser(u.Name, u.Email)

	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"name":     u.Name,
		"email":    u.Email,
		"database": "user",
	})
}
