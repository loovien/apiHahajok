/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/16 1:17
 * @description: 
 */

package service

import "github.com/vvotm/apiHahajok/models"

func RecordOpenID(openId, unionId string) (int64, error) {
	usermodel := models.NewUser()
	return usermodel.RecordOpenId(openId, unionId)
}
