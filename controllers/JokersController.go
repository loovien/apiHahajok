/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/16 13:42
 * @description: 
 */

package controllers

import (
	"github.com/labstack/echo"
	"github.com/vvotm/apiHahajok/models/request"
	"github.com/vvotm/apiHahajok/service"
)

func GetLatestsJokersList(ctx echo.Context) (err error) {
	pageInfo := new(request.ReqPage)
	if err = ctx.Bind(pageInfo); err != nil {
		return err
	}
	service.GetLatestJokersList(pageInfo)
}
