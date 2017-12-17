/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/16 13:10
 * @description: 
 */

package dao

import (
	"testing"
	"github.com/vvotm/apiHahajok/db"
	"github.com/vvotm/apiHahajok/models/request"
)

func TestUser_RecordOpenId(t *testing.T) {
	var a int64 = 1
	t.Log(a)
	t.Log(int(a))
	return
	userDao := NewUser()
	dbConn := db.GetConn()
	dbConn = dbConn.Create(userDao)
	t.Log(dbConn.Error)
	t.Log(userDao.Id)
	return
	id, err := userDao.RecordOpenId("abcdefg", "你没啊")
	t.Log(id)
	t.Log(err)
}

func TestUser_UpdateUserInfo(t *testing.T) {
	userRepo := NewUser()
	r := request.ReqUserInfo{
		OpenId: "111",
		UnionId: "xxxxxxxxxxxx",
	}
	t.Log(userRepo.UpdateUserInfo(&r))

}

func TestUser_GetUserInfoById(t *testing.T) {
	userDao := NewUser()
	resultSet := userDao.GetUserInfoById(1)
	t.Log(resultSet)

}
