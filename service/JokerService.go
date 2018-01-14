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
	"github.com/vvotm/apiHahajok/dao/criteria"
)

func GetLatestJokersList(pageInfo *request.ReqPage) (respJokerList response.RespJokerList, err error) {
	respJokerList.Size = pageInfo.Size

	jokerDao := dao.NewJoker()
	commonCriteria := criteria.CommonCriteria{
		Condition: "status = 1",
	}
	respJokerList.Cnt = jokerDao.Count(commonCriteria)
	jokerList, err := jokerDao.GetJokerList(criteria.PageCriteria{
		Page: pageInfo.Page,
		Size: pageInfo.Size,
		CommonCriteria: commonCriteria,

	})
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
	return respJokerList, nil
}

func GetHotsJokersList(pageInfo *request.ReqPage) (respJokerList response.RespJokerList, err error) {
	respJokerList.Size = pageInfo.Size

	jokerDao := dao.NewJoker()
	commonCriteria := criteria.CommonCriteria{
		Condition: "status = 1",
		Order: "replies DESC",
	}
	respJokerList.Cnt = jokerDao.Count(commonCriteria)
	jokerList, err := jokerDao.GetJokerList(criteria.PageCriteria{
		Page: pageInfo.Page,
		Size: pageInfo.Size,
		CommonCriteria: commonCriteria,

	})
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
	return respJokerList, nil
}

func GetJokersById(id int) (*response.RespJokers, error) {
	jokerDao := dao.NewJoker()
	query := criteria.CommonCriteria{
		Condition: "id = ?",
		ConditionBind: []interface{}{id},
	}
	joker, _ := jokerDao.GetOne(query)

	userDao := dao.NewUser()
	member := userDao.GetUserInfoById(joker.Uid)
	classificationDao := dao.NewClassification()
	classification := classificationDao.GetClassificationById(joker.ClassId)
	return &response.RespJokers{
		Joker: *joker,
		Member: response.RespUser{
			User: member,
		},
		Classification: response.RespClassification{
			Classification: classification,
		},
	}, nil
}
