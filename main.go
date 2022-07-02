package main

import (
	"crypto/subtle"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	. "test/mtemplate"
	. "test/router"
)

type (
	User struct {
		Name  string `json:"name" validate:"required" form:"name"`
		Email string `json:"email" validate:"required,email" form:"email"`
	}

	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {

		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.StructField())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println("---------------")
		}

		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Renderer = Rend()
	g := e.Group("/admin")

	Routes(e)
	e.POST("/users", func(c echo.Context) (err error) {
		u := new(User)
		if err = c.Bind(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err = c.Validate(u); err != nil {
			return err
		}
		signingKey := []byte("secret")
		// return c.JSON(http.StatusOK, u)
		return c.Render(http.StatusOK, "home.html", map[string]interface{}{
			"user": u,
		})
	})
	e.Logger.Fatal(e.Start(":1323"))
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// Be careful to use constant time comparison to prevent timing attacks
		if subtle.ConstantTimeCompare([]byte(username), []byte("joe")) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte("secret")) == 1 {
			return true, nil

		}
		return false, nil
	}))
	g.GET("/here", Home)
	e.Use(middleware.BasicAuthWithConfig(middleware.BasicAuthConfig{}))

	config := middleware.JWTConfig{
		TokenLookup: "query:token",
		ParseTokenFunc: func(auth string, c echo.Context) (interface{}, error) {
			keyFunc := func(t *jwt.Token) (interface{}, error) {
				if t.Method.Alg() != "HS256" {
					return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
				}
				return signingKey, nil
			}

			// claims are of type `jwt.MapClaims` when token is created with `jwt.Parse`
			token, err := jwt.Parse(auth, keyFunc)
			if err != nil {
				return nil, err
			}
			if !token.Valid {
				return nil, err
			}
			return token, nil
		},
	}

	g.Use(middleware.JWTWithConfig(config))

	e.Static("/static", "static")
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	e.Logger.Fatal(e.Start(":1323"))
}
func Home(c echo.Context) error {

	return c.Render(http.StatusOK, "home.html", map[string]interface{}{})
}
