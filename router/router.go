package router

import (
	. "test/handler/get/home"
	. "test/handler/post/form"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {

	e.GET("/", Home)

	e.POST("/form", Form)
}
