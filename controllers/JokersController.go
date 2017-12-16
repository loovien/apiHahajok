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
	"github.com/vvotm/apiHahajok/errhandle"
	"net/http"
	"github.com/vvotm/apiHahajok/utils"
)

func GetLatestsJokersList(ctx echo.Context) (err error) {
	pageInfo := request.NewReqPage()
	if err = ctx.Bind(pageInfo); err != nil {
		return err
	}
	resp, err := service.GetLatestJokersList(pageInfo)
	errhandle.CheckError(err)
	return ctx.JSON(http.StatusOK, utils.GetCommonResp(resp, errhandle.SUCCESS_CODE, "success"))
}
