/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/16 0:35
 * @description: 
 */

package boot

import (
	"github.com/labstack/echo"
	"sync"
	"github.com/vvotm/apiHahajok/middleware"
	"github.com/vvotm/apiHahajok/routes"
)

var (
	once *sync.Once = &sync.Once{}
	appInst *App = nil
)

type App struct {
	AppInst *echo.Echo
}

func GetApp() *App {
	if appInst != nil {
		return appInst
	}
	once.Do(func() {
		appInst = &App{
			AppInst: newEchoApp(),
		}
	})
	return appInst
}

func newEchoApp() *echo.Echo {
	app := echo.New()
	middleware.InitMiddleware(app) // initialize middleware
	routes.InitRoutes(app) // initialize route
	return app
}
