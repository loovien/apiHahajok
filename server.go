package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	webapp := echo.New()
	webapp.GET("/webapp", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world")
	})
	webapp.Start(":8080")
}
