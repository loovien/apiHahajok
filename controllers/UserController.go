package controllers

import (
	"github.com/labstack/echo"
	"strings"
	"net/http"
	"github.com/vvotm/apiHahajok/utils"
	"github.com/vvotm/apiHahajok/errhandle"
)

func GetUserOpenID(ctx echo.Context) error {
	jsCode := strings.Trim(ctx.FormValue("jscode"), " ")
	if jsCode == "" {
		return ctx.JSON(http.StatusBadRequest, utils.GetCommonResp(nil, errhandle.ERROR_CODE, "jscode can't empty!"))
	}
	userOpenId, err := utils.GetOpenID(jsCode)
	if err != nil {
		return ctx.JSON(http.StatusOK, utils.GetCommonResp(nil, errhandle.ERROR_CODE, err.Error()))
	}
	if errorCode := userOpenId["errcode"]; errorCode != nil {
		return  ctx.JSON(http.StatusOK, utils.GetCommonResp("", int(errorCode.(float64)), "network busing!"))
	}
	return ctx.JSON(http.StatusOK, utils.GetCommonResp(userOpenId, errhandle.SUCCESS_CODE, "success"))
}
