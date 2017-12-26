/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/16 18:02
 * @description: 
 */

package dao

import (
	"github.com/vvotm/apiHahajok/dao/criteria"
	"github.com/vvotm/apiHahajok/db"
	"github.com/vvotm/apiHahajok/errhandle"
)

const (
	STATUS_DELETE int = -1
	STATUS_CHECKING int = 0
	STATUS_OK int = 1
)

type Replies struct {
	Id int `json:"id" gorm:"primary_key"`
	JokerId int `json:"jokerId" gorm:"column:jokerId"`
	Uid int `json:"uid"`
	Content string `json:"content"`
	ImageList string `json:"imageList" gorm:"column:imageList"`
	Status int `json:"status"`
	CreatedAt int `json:"createdAt" gorm:"column:createdAt"`
	PassedAt int `json:"passedAt" gorm:"column:passedAt"`
	UpdatedAt int `json:"updatedAt" gorm:"column:updatedAt"`
}

func NewReplies() *Replies {
	return &Replies{}
}

func (r *Replies) TableName () string  {
	return "replies"
}

func (r *Replies) Count(criteria criteria.CommonCriteria) (cnt int) {
	dbConn := db.GetConn().Model(r)
	dbConn = ApplyCommonQuery(dbConn, criteria)
	dbConn.Count(&cnt)
	return cnt
}

func (r *Replies) GetRepliesList(criteria criteria.PageCriteria) (repliesList []Replies, err error) {
	dbConn := db.GetConn().Model(r)
	dbConn = ApplyPageQuery(dbConn, criteria)

	rows, err := dbConn.Rows()
	defer rows.Close()
	if err != nil {
		return nil, errhandle.NewPDOError("查询数据失败", errhandle.DB_OPERATE_ERROR, err.Error())
	}

	for rows.Next() {
		replies := Replies{}
		dbConn.ScanRows(rows, &replies)
		repliesList = append(repliesList, replies)
	}
	return repliesList, nil
}

func (r *Replies) Insert() int {
	dbConn := db.GetConn().Debug().Model(r)
	dbConn.Save(r)
	return r.Id
}
