/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/26 20:13
 * @description: 
 */

package controllers

import (
	"github.com/labstack/echo"
	"github.com/vvotm/apiHahajok/models/request"
	"strconv"
	"github.com/vvotm/apiHahajok/service"
	"github.com/vvotm/apiHahajok/errhandle"
	"net/http"
	"github.com/vvotm/apiHahajok/utils"
	"github.com/vvotm/apiHahajok/models/validate"
)

func GetRepliesListByJokerId(ctx echo.Context) (err error)  {
	pageInfo := request.NewReqPage()
	if err = ctx.Bind(pageInfo); err != nil {
		return  err;
	}
	jokerId, _ := strconv.Atoi(ctx.Param("id"))

	resp, err := service.GetRepliesByJokerId(pageInfo, jokerId)
	errhandle.CheckError(err)
	return ctx.JSON(http.StatusOK, utils.GetCommonResp(resp, errhandle.SUCCESS_CODE, "success"))
}

func PostReplies(ctx echo.Context) (err error)  {
	replies := request.NewReqReplies()
	if err = ctx.Bind(replies); err != nil {
		return  err;
	}
	repliesValidate := validate.NewRepliesValidate()
	if err = repliesValidate.Validate(replies); err != nil {
		return err
	}
	lastInsertId, err := service.PostReplies(replies)
	if err != nil {
		return err
	}
	resp := struct {
		LastId int `json:"lastId"`
	}{LastId: lastInsertId}
	return ctx.JSON(http.StatusOK, utils.GetCommonResp(resp, errhandle.SUCCESS_CODE, "success"))
}
