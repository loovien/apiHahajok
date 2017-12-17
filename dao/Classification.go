/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/16 18:03
 * @description: 
 */

package dao

import "github.com/vvotm/apiHahajok/db"

type Classification struct {
	Id  int `json:"id"`
	Name  string `json:"name"`
	Icon string `json:"icon"`
	CreatedAt int `json:"createdAt"`
	UpdatedAt int `json:"updatedAt"`
}

func NewClassification() *Classification {
	return &Classification{}
}

func (c *Classification) GetClassificationById(id int) (classification Classification) {
	dbConn := db.GetConn()
	sql := "select * from classification where id = ?"
	dbConn.Raw(sql, id).Scan(&classification)
	return classification
}
