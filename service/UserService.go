/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/16 1:17
 * @description: 
 */

package service

import (
	"github.com/vvotm/apiHahajok/models"
	"errors"
	"github.com/labstack/gommon/log"
)

func RecordOpenID(openId, unionId string) (int64, error) {
	usermodel := models.NewUser()
	lastInsertId, err := usermodel.RecordOpenId(openId, unionId)
	if err != nil {
		log.Errorf(err)
		return 0, errors.New("系统繁忙, 请稍候再试!")
	}
	return lastInsertId, nil
}
