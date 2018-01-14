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
	"strconv"
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

func GetHotsJokersList(ctx echo.Context) (err error) {
	pageInfo := request.NewReqPage()
	if err = ctx.Bind(pageInfo); err != nil {
		return err
	}
	resp, err := service.GetHotsJokersList(pageInfo)
	errhandle.CheckError(err)
	return ctx.JSON(http.StatusOK, utils.GetCommonResp(resp, errhandle.SUCCESS_CODE, "success"))
}

func GetJokersById(ctx echo.Context) (err error)  {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if id <= 0 {
		return ctx.JSON(http.StatusNotFound, utils.GetCommonResp(nil, errhandle.ERROR_CODE, "参数错误"));
	}
	resp, err := service.GetJokersById(id)
	return ctx.JSON(http.StatusOK, utils.GetCommonResp(resp, errhandle.SUCCESS_CODE, "success"))
}

func GetClassJokersList(ctx echo.Context) (err error) {
	pageInfo := request.NewReqPage()
	if err = ctx.Bind(pageInfo); err != nil {
		return err
	}
	classId, _ := strconv.Atoi(ctx.Param("classId"))
	if classId <= 0 {
		return ctx.JSON(http.StatusOK, utils.GetCommonResp(nil, errhandle.ERROR_CODE, "参数错误"))
	}

	resp, err := service.GetClassJokersList(pageInfo, classId)
	errhandle.CheckError(err)

	return ctx.JSON(http.StatusOK, utils.GetCommonResp(resp, errhandle.SUCCESS_CODE, "success"))
}
