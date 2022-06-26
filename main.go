package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	. "test/mtemplate"
	. "test/router"
)

func main() {
	e := echo.New()

	e.Renderer = Rend()

	Routes(e)
	e.Static("/static", "static")
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	e.Logger.Fatal(e.Start(":1323"))
}
