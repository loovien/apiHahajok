/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/17 21:52
 * @description: 
 */

package controllers

import (
	"github.com/labstack/echo"
	"github.com/vvotm/apiHahajok/models/request"
	"github.com/vvotm/apiHahajok/service"
	"net/http"
	"github.com/vvotm/apiHahajok/utils"
	"github.com/vvotm/apiHahajok/errhandle"
)

func GetClassificationList(ctx echo.Context) (err error)  {
	pageInfo := request.NewReqPage()
	if err = ctx.Bind(pageInfo); err != nil {
		return err
	}
	resp, err := service.GetClassificationList(pageInfo)
	return ctx.JSON(http.StatusOK, utils.GetCommonResp(resp, errhandle.SUCCESS_CODE, "success"))
}
