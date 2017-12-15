/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/15 23:42
 * @description: webapp initialize middleware
 */

package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func InitMiddleware(e *echo.Echo)  {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{ // set cors configure
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
	}))
}
