/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/16 1:17
 * @description: 
 */

package service

import (
	"github.com/vvotm/apiHahajok/models"
	"github.com/vvotm/apiHahajok/errhandle"
)

func RecordOpenID(openId, unionId string) (int64, error) {
	usermodel := models.NewUser()
	lastInsertId, err := usermodel.RecordOpenId(openId, unionId)
	if err != nil {
		return 0, errhandle.NewPDOError("数据库操作失败:" + err.Error())
	}
	return lastInsertId, nil
}
