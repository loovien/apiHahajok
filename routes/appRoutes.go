package routes

import (
	"github.com/labstack/echo"
	"github.com/vvotm/apiHahajok/controllers"
)

func InitRoutes(e echo.Echo) {
	e.POST("/user/openid.json", controllers.GetUserOpenID)
}
