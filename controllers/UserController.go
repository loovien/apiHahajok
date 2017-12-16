package controllers

import (
	"github.com/labstack/echo"
	"strings"
	"net/http"
	"github.com/vvotm/apiHahajok/utils"
	"github.com/vvotm/apiHahajok/errhandle"
	"github.com/vvotm/apiHahajok/service"
	"errors"
	"github.com/vvotm/apiHahajok/models/request"
	"github.com/labstack/gommon/log"
)

func GetUserOpenID(ctx echo.Context) error {
	jsCode := strings.Trim(ctx.FormValue("jscode"), " ")
	if jsCode == "" {
		return ctx.JSON(http.StatusBadRequest, utils.GetCommonResp(nil, errhandle.ERROR_CODE, "参数jscode缺失!"))
	}
	userOpenId, err := utils.GetOpenID(jsCode)
	if err != nil {
		return ctx.JSON(http.StatusOK, utils.GetCommonResp(nil, errhandle.ERROR_CODE, err.Error()))
	}
	if errorCode := userOpenId["errcode"]; errorCode != nil {
		return  ctx.JSON(http.StatusOK, utils.GetCommonResp(userOpenId, int(errorCode.(float64)), "抱歉, 网络繁忙!"))
	}
	go func() {
		var openId, unionId string
		if userOpenId["openid"] != nil {
			openId = userOpenId["openid"].(string)
		}
		if userOpenId["unionid"] != nil {
			unionId = userOpenId["unionid"].(string)
		}
		ctx.Echo().Logger.Infof("Info: save: %v", userOpenId)
		if openId != "" {
			service.RecordOpenID(openId, unionId)
		}
	}()
	ctx.Echo().Logger.Infof("Info: UserOpenId: %v", userOpenId)
	return ctx.JSON(http.StatusOK, utils.GetCommonResp(userOpenId, errhandle.SUCCESS_CODE, "success"))
}

func UpdateUserInfo(ctx echo.Context) (err error) {
	reqUserInfo := new(request.ReqUserInfo)
	if err = ctx.Bind(reqUserInfo); err != nil {
		return errors.New("参数缺失:" + err.Error())
	}
	err = service.UpdateUserInfo(reqUserInfo)
	if err != nil {
		log.Info(err)
		return ctx.JSON(http.StatusOK, utils.GetCommonResp(nil, errhandle.ERROR_CODE, err.Error()))
	}
	return ctx.JSON(http.StatusOK, utils.GetCommonResp(reqUserInfo, errhandle.SUCCESS_CODE, "success"))
}
