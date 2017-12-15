/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/16 1:45
 * @description: 
 */

package models

import (
	"github.com/vvotm/apiHahajok/db"
	"time"
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
	sql := "insert into user (openId, unionId, createdAt, updatedAt) values (?, ?, ?, ?)"
	result, err := dbConn.Exec(sql, openId, unionId, nowTime, nowTime)
	if err != nil{
		return 0, err
	}
	return result.LastInsertId()
}

