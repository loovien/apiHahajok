package routes

import (
	"github.com/labstack/echo"
	"github.com/vvotm/apiHahajok/controllers"
)

func InitRoutes(e *echo.Echo) {
	e.POST("/user/openid.json", controllers.GetUserOpenID)
	e.PUT("/user/infos.json", controllers.UpdateUserInfo)

	e.GET("/jokers/latests.json", controllers.GetLatestsJokersList)
	e.GET("/jokers/:id/info.json", controllers.GetJokersById)
	e.GET("/jokers/hots.json", controllers.GetHotsJokersList)
	e.GET("/jokers/class/:classId/list.json", controllers.GetClassJokersList)


	e.GET("/classification/list.json", controllers.GetClassificationList)

	e.POST("/replies.json", controllers.PostReplies)
	e.GET("/replies/jokers/:id/list.json", controllers.GetRepliesListByJokerId)

}
