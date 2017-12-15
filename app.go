package main

import (
	"github.com/labstack/echo"
	"github.com/vvotm/apiHahajok/routes"
	"github.com/vvotm/apiHahajok/middleware"
)

func main() {
	app := echo.New()
	middleware.InitMiddleware(app) // initialize middleware
	routes.InitRoutes(app) // initialize route

	err := app.Start(":8080")
	app.Logger.Errorf("系统错误: %v", err)
}
