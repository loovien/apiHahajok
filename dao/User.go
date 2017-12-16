/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/16 1:45
 * @description: 
 */

package dao

import (
	"github.com/vvotm/apiHahajok/db"
	"github.com/vvotm/apiHahajok/errhandle"
	"time"
	"github.com/vvotm/apiHahajok/models/request"
)

const (
	NO_SAVE = 0
	IS_SAVE = iota
)
type User struct {

}

func NewUser() *User {
	return &User{}
}

func (u *User) RecordOpenId(openId, unionId string) (int64, error) {
	dbConn := db.GetConn()
	nowTime := time.Now().Unix()
	countSql := "select id from user where openId = ?"
	var count int64 = 0
	dbConn.QueryRow(countSql, openId).Scan(&count)
	if count > 0 {
		return count, nil
	}
	sql := "insert into user (openId, unionId, issave, createdAt, updatedAt) values (?, ?, ?, ?, ?)"
	result, err := dbConn.Exec(sql, openId, unionId, IS_SAVE, nowTime, nowTime)
	if err != nil {
		return 0, errhandle.NewPDOError("数据插入失败: " + err.Error(), errhandle.DB_OPERATE_ERROR)
	}
	return result.LastInsertId()
}

func (u *User) UpdateUserInfo(r *request.ReqUserInfo) (err error) {
	dbConn := db.GetConn()
	nowTime := time.Now().Unix()
	countSql := "select id from user where openId = ? and issave = ? and updatedAt > ?"
	var count int64 = 0
	// 10 minute don't update
	dbConn.QueryRow(countSql, r.OpenId, IS_SAVE, time.Now().Unix() - 600).Scan(&count)
	if count > 0 {
		return  nil
	}
	updateSql := `update user unionId = ?, nickname = ?, avatar = ?, gender = ?, country = ?
		 province = ?, city = ?, lang = ?, updatedAt = ? where openId = ?`
	_, err = dbConn.Exec(updateSql, r.UnionId, r.Nickname, r.Avatar, r.Gender, r.Country, r.Province, r.City,
		r.Lang, nowTime, r.OpenId)
	return  err
}
