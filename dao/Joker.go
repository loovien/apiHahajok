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
	"fmt"
)

type Joker struct {
	Id int `json:"id"`
	Uid int `json:"uid"`
	ClassId int `json:"classId"`
	Title string `json:"title"`
	Content string `json:"content"`
	ImageList string `json:"imageList"`
	MediaUrl string `json:"mediaUrl"`
	Replies int `json:"replies"`
	Status int `json:"status"`
	CreatedAt int `json:"createdAt"`
	PassedAt int `json:"passedAt"`
	UpdatedAt int `json:"updatedAt"`
}

func NewJoker() *Joker {
	return &Joker{}
}

func (j *Joker) Count(conditions string) (cnt int) {
	dbConn := db.GetConn()
	sql := "select count(*) as cnt from joker where ?";
	if conditions == "" {
		conditions = "1 = 1";
	}
	dbConn.QueryRow(sql, conditions).Scan(&cnt)
	return cnt
}

func (j *Joker) GetJokerList(columns, conditions string) (jokerList[]Joker, err error) {
	dbConn := db.GetConn()
	if conditions == "" {
		conditions = "1 = 1"
	}
	sql := fmt.Sprintf("select %s from joker where %s", columns, conditions)
	log.Info(sql)
	rows, err := dbConn.Query(sql)
	defer rows.Close()
	if err != nil {
		log.Error(err)
		return nil, errhandle.NewPDOError("查询数据失败", errhandle.DB_OPERATE_ERROR, err.Error())
	}
	var columnsSlice, _ = rows.Columns()
	jokerList = []Joker{}
	for rows.Next()  {
		var joker = Joker{}
		scanColumnList := j.getScanColumnList(columnsSlice, &joker)
		rows.Scan(scanColumnList...)
		jokerList = append(jokerList, joker)
	}
	return jokerList, nil
}

func (j *Joker) DeleteJoker(conditions string) (err error) {
	dbConn := db.GetConn()
	sql := "delete from jocker where ?"
	_, err = dbConn.Exec(sql, conditions)
	return errhandle.NewPDOError("删除失败", errhandle.DB_OPERATE_ERROR, err.Error())
}

func (j *Joker) DeleteJokerById(id int) (err error) {
	dbConn := db.GetConn()
	sql := "delete from jocker where id = ?"
	_, err = dbConn.Exec(sql, id)
	return errhandle.NewPDOError("删除失败", errhandle.DB_OPERATE_ERROR, err.Error())
}

func (j *Joker) getScanColumnList(columns []string, joker *Joker) []interface{} {
	var scanColumnList = []interface{}{}
	for _, column := range columns {
		switch column {
		case "id":
			scanColumnList = append(scanColumnList, &joker.Id)
		case "uid":
			scanColumnList = append(scanColumnList, &joker.Uid)
		case "classId":
			scanColumnList = append(scanColumnList, &joker.ClassId)
		case "title":
			scanColumnList = append(scanColumnList, &joker.Title)
		case "content":
			scanColumnList = append(scanColumnList, &joker.Content)
		case "imageList":
			scanColumnList = append(scanColumnList, &joker.ImageList)
		case "mediaUrl":
			scanColumnList = append(scanColumnList, &joker.MediaUrl)
		case "replies":
			scanColumnList = append(scanColumnList, &joker.Replies)
		case "status":
			scanColumnList = append(scanColumnList, &joker.Status)
		case "createdAt":
			scanColumnList = append(scanColumnList, &joker.CreatedAt)
		case "passedAt":
			scanColumnList = append(scanColumnList, &joker.PassedAt)
		case "updatedAt":
			scanColumnList = append(scanColumnList, &joker.UpdatedAt)
		}
	}
	return scanColumnList
}
