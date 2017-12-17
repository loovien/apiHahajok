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
	"github.com/labstack/gommon/log"
)

const (
	NO_SAVE = 0
	IS_SAVE = iota
)
type User struct {
	Id int `json:"id" gorm:"primary_key"`
	OpenId string `json:"openId" gorm:"column:openId"`
	UnionId string `json:"unionId" gorm:"column:unionId"`
	Nickname string `json:"nickname" gorm:"nickname"`
	Wallet float64 `json:"wallet"`
	Avatar string `json:"avatar"`
	Gender int `json:"gender"`
	Country string `json:"country"`
	Province string `json:"province"`
	City string `json:"city"`
	Lang string `json:"lang"`
	Issave int `json:"issave"`
	CreatedAt int `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt int `json:"updatedAt" gorm:"column:updatedAt"`
}

func (User) TableName() string {
	return "user"
}

func NewUser() *User {
	return &User{}
}

func (u *User) RecordOpenId(openId, unionId string) (int, error) {
	dbConn := db.GetConn()

	var count int
	dbConn.Model(u).First("openId = ?", openId).Count(&count)
	if count > 0 {
		return 0, nil
	}
	nowTime := int(time.Now().Unix())
	userRow := User{OpenId:openId, UnionId:unionId, Issave:IS_SAVE, CreatedAt:nowTime, UpdatedAt:nowTime}
	dbConn = dbConn.Create(userRow)
	if dbConn.Error != nil {
		log.Error(dbConn.Error)
		return 0, errhandle.NewPDOError("插入数据库失败", errhandle.DB_OPERATE_ERROR, dbConn.Error.Error())
	}
	return userRow.Id, nil
}

func (u *User) UpdateUserInfo(r *request.ReqUserInfo) (err error) {
	dbConn := db.GetConn()
	nowTime := time.Now().Unix()
	countSql := "select id from user where openId = ? and issave = ? and updatedAt > ?"
	var count int = 0
	// 10 minute don't update
	dbConn.Raw(countSql, r.OpenId, IS_SAVE, time.Now().Unix() - 600).Scan(&count)
	log.Info(count)
	if count > 0 {
		return  nil
	}
	updateSql := "update user set unionId = ?, nickname = ?, avatar = ?, gender = ?, country = ?," +
		 "province = ?, city = ?, lang = ?, issave = ?, updatedAt = ? where openId = ?"
	log.Infof("SQL:%s DATA: %v", updateSql, r)
	dbConn = dbConn.Exec(updateSql, r.UnionId, r.Nickname, r.Avatar, r.Gender, r.Country, r.Province, r.City,
		r.Lang, IS_SAVE, nowTime, r.OpenId)
	if dbConn.Error != nil {
		log.Error(err)
		return  errhandle.NewPDOError("更新数据失败!", errhandle.DB_OPERATE_ERROR, dbConn.Error.Error())
	}
	return nil
}

func (u *User) GetUserInfoById(id int) (userInfo User) {
	dbConn := db.GetConn()
	dbConn.First(&userInfo, id)
	return userInfo
}

