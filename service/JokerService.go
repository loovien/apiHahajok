/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/16 13:44
 * @description: 
 */

package service

import (
	"github.com/vvotm/apiHahajok/models/request"
	"github.com/vvotm/apiHahajok/models/response"
	"github.com/vvotm/apiHahajok/dao"
	"fmt"
)

func GetLatestJokersList(pageInfo *request.ReqPage) (respJokerList response.RespJokerList, err error) {
	respJokerList.Size = pageInfo.Size

	jokerDao := dao.NewJoker()
	respJokerList.Cnt = jokerDao.Count("status = 1")
	conditions := fmt.Sprintf("status = 1 limit %d, %d", (pageInfo.Page - 1) * pageInfo.Size, pageInfo.Size )
	jokerList, err := jokerDao.GetJokerList("*", conditions)
	if err != nil {
		return respJokerList, err
	}
	userDao := dao.NewUser()
	classificationDao := dao.NewClassification()
	for _, joker := range jokerList{
		member := response.RespUser{User:userDao.GetUserInfoById(joker.Uid)}
		classification := response.RespClassification{
			Classification:classificationDao.GetClassificationById(joker.ClassId)}
		respJoker := response.RespJokers{Joker:joker, Member:member, Classification:classification}
		respJokerList.List = append(respJokerList.List, respJoker)
	}
	fmt.Println(respJokerList)
	return respJokerList, nil
}
