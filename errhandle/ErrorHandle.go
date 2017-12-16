/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/16 12:01
 * @description: 
 */

package errhandle

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/vvotm/apiHahajok/utils"
	"github.com/labstack/gommon/log"
)

func ErrorHandle(err error, ctx echo.Context)  {
	defer func() {
		if r := recover(); r != nil {
			ctx.JSON(http.StatusOK, utils.GetCommonResp(struct {
				ErrMsg string `json:"errmsg"`
			}{err.Error()}, ERROR_CODE, "抱歉, 网络繁忙, 请稍候再试!"))
		}
	}()
	webErr := err.(AbsErrer)
	ctx.JSON(http.StatusOK, utils.GetCommonResp(struct{
		ErrMsg string `json:"errmsg"`
	}{webErr.ErrorMsg()}, webErr.GetCode(), webErr.Error()))
}

func CheckError(err error)  {
	if err != nil {
		log.Error(err)
		panic(err)
	}
}
