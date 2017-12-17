/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/16 18:03
 * @description: 
 */

package dao

import (
	"github.com/vvotm/apiHahajok/db"
	"github.com/vvotm/apiHahajok/dao/criteria"
	"github.com/vvotm/apiHahajok/errhandle"
)

type Classification struct {
	Id  int `json:"id"`
	Name  string `json:"name"`
	Icon string `json:"icon"`
	Status int `json:"status"`
	CreatedAt int `json:"createdAt"`
	UpdatedAt int `json:"updatedAt"`
}

func (Classification) TableName() string  {
	return "classification"
}

func NewClassification() *Classification {
	return &Classification{}
}

func (c *Classification) GetClassificationById(id int) (classification Classification) {
	dbConn := db.GetConn().Model(c)
	dbConn.First(&classification, id)
	return classification
}

func (c *Classification) Count(criteria criteria.CommonCriteria) (cnt int)  {
	dbConn := db.GetConn().Model(c)
	dbConn = ApplyCommonQuery(dbConn, criteria)
	dbConn.Count(&cnt)
	return cnt
}

func (c *Classification) GetClassificationList(criteria criteria.PageCriteria) (classificationList []Classification, err error)  {
	dbConn := db.GetConn().Model(c)
	dbConn = ApplyPageQuery(dbConn, criteria)

	rows, err := dbConn.Rows()
	defer rows.Close()
	if err != nil {
		return classificationList, errhandle.NewPDOError("查询数据错误", errhandle.DB_OPERATE_ERROR, err.Error())
	}
	for rows.Next() {
		classification := Classification{}
		dbConn.ScanRows(rows, &classification)
		classificationList = append(classificationList, classification)
	}
	return classificationList, nil
}