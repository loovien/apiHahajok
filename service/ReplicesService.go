/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/26 20:17
 * @description: 
 */

package service

import (
	"github.com/vvotm/apiHahajok/models/request"
	"github.com/vvotm/apiHahajok/models/response"
	"github.com/vvotm/apiHahajok/dao"
	"github.com/vvotm/apiHahajok/dao/criteria"
)

func GetRepliesByJokerId(pageInfo *request.ReqPage, id int) (respRepliesList response.RespRepliesList, err error)  {
	repliesDao := dao.NewReplies()
	commonCriteria := criteria.CommonCriteria{
		Condition: "jokerId = ?",
		ConditionBind: []interface{}{id},
		Order: "id DESC",
	}
	respRepliesList.Cnt = repliesDao.Count(commonCriteria)
	respRepliesList.Size = pageInfo.Size

	pageCriteria := criteria.PageCriteria{
		Page: pageInfo.Page,
		Size: pageInfo.Size,
		CommonCriteria: commonCriteria,
	}

	repliesList, err := repliesDao.GetRepliesList(pageCriteria)
	if err != nil {
		return respRepliesList, err
	}
	for _, replies := range repliesList {
		respReplies := response.RespReplies{Replies: replies}
		respRepliesList.List = append(respRepliesList.List, respReplies)
	}
	return respRepliesList, nil
}
