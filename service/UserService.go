/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/16 1:17
 * @description: 
 */

package service

import (
	"errors"
	"github.com/labstack/gommon/log"
	"github.com/vvotm/apiHahajok/dao"
	"github.com/vvotm/apiHahajok/models/request"
)

func RecordOpenID(openId, unionId string) (int, error) {
	userDao := dao.NewUser()
	lastInsertId, err := userDao.RecordOpenId(openId, unionId)
	if err != nil {
		log.Error(err)
		return 0, errors.New("系统繁忙, 请稍候再试!")
	}
	return lastInsertId, nil
}

func UpdateUserInfo(r *request.ReqUserInfo) error {
	userDao := dao.NewUser()
	return userDao.UpdateUserInfo(r)
}