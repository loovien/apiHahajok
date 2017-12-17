/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/16 18:02
 * @description: 
 */

package dao

import (
	"github.com/vvotm/apiHahajok/db"
	"github.com/vvotm/apiHahajok/errhandle"
	"github.com/labstack/gommon/log"
	"github.com/vvotm/apiHahajok/dao/criteria"
)

type Joker struct {
	Id int `json:"id" gorm:"primary_key"`
	Uid int `json:"uid"`
	ClassId int `json:"classId" gorm:"column:classId"`
	Title string `json:"title"`
	Content string `json:"content"`
	ImageList string `json:"imageList" gorm:"column:imageList"`
	MediaUrl string `json:"mediaUrl" gorm:"column:mediaUrl"`
	Replies int `json:"replies"`
	Status int `json:"status"`
	CreatedAt int `json:"createdAt" gorm:"column:createdAt"`
	PassedAt int `json:"passedAt" gorm:"column:passedAt"`
	UpdatedAt int `json:"updatedAt" gorm:"column:updatedAt"`
}

func (Joker) TableName() string {
	return "joker"
}

func NewJoker() *Joker {
	return &Joker{}
}

func (j *Joker) Count(criteria criteria.CommonCriteria) (cnt int) {
	dbConn := db.GetConn().Model(j)
	dbConn = ApplyCommonQuery(dbConn, criteria)
	dbConn.Count(&cnt)
	return cnt
}

func (j *Joker) GetJokerList(criteria criteria.PageCriteria) (jokerList[]Joker, err error) {
	dbConn := db.GetConn().Debug().Model(j)
	dbConn = ApplyPageQuery(dbConn, criteria)
	rows, err := dbConn.Rows()
	defer rows.Close()
	if err != nil {
		log.Error(err)
		return nil, errhandle.NewPDOError("查询数据失败", errhandle.DB_OPERATE_ERROR, err.Error())
	}
	for rows.Next()  {
		var joker = Joker{}
		dbConn.ScanRows(rows, &joker)
		jokerList = append(jokerList, joker)
	}
	return jokerList, nil
}

func (j *Joker) DeleteJoker(conditions string, bind ...interface{}) (err error) {
	dbConn := db.GetConn().Model(j)
	dbConn.Where(conditions, bind...).Delete(j)
	if dbConn.Error != nil {
		return errhandle.NewPDOError("删除失败", errhandle.DB_OPERATE_ERROR, dbConn.Error.Error())
	}
	return nil
}

func (j *Joker) DeleteJokerById(id int) (err error) {
	dbConn := db.GetConn()
	sql := "delete from jocker where id = ?"
	dbConn.Exec(sql, id)
	if dbConn.Error != nil {
		return errhandle.NewPDOError("删除失败", errhandle.DB_OPERATE_ERROR, dbConn.Error.Error())
	}
	return nil
}

