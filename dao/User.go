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
	var nowTime int64 = time.Now().Unix()
	dbConn := db.GetConn()
	userRepo := User{}
	// 10 minute don't update
	log.Info(r)
	dbConn.Model(&userRepo).Where("openId = ? and issave = ? and updatedAt < ?",
		r.OpenId, IS_SAVE, time.Now().Unix() - 600).First(&userRepo)
	if userRepo.Id <= 0 {
		return  nil
	}
	updateUserDao := User{
		Id: userRepo.Id,
		OpenId: userRepo.OpenId,
		UnionId: r.UnionId,
		Nickname: r.Nickname,
		Avatar: r.Avatar,
		Gender: r.Gender,
		Country: r.Country,
		Province: r.Province,
		City: r.City,
		Lang: r.Lang,
		Issave: IS_SAVE,
		UpdatedAt: int(nowTime),
	}
	dbConn = dbConn.Save(&updateUserDao)
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

