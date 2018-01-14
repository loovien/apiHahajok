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
	"time"
	"errors"
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
		userDao := dao.NewUser()
		userInfo := userDao.GetUserInfoById(replies.Uid)
		respReplies := response.RespReplies{
			Replies: replies,
			Member: userInfo,
		}
		respRepliesList.List = append(respRepliesList.List, respReplies)
	}
	return respRepliesList, nil
}

func PostReplies(reqReplies *request.ReqReplies) (lastInsertId int, err error) {
	repliesDao := dao.NewReplies()
	nowTime := int(time.Now().Unix())

	repliesDao.Uid = reqReplies.Uid
	repliesDao.JokerId = reqReplies.JokerId
	repliesDao.Content = reqReplies.Content
	repliesDao.ImageList = reqReplies.ImageList
	repliesDao.Status = dao.STATUS_CHECKING
	repliesDao.CreatedAt = nowTime
	repliesDao.UpdatedAt = nowTime
	lastInsertId = repliesDao.Insert()
	if lastInsertId == 0 {
		return 0, errors.New("插入数据库失败")
	}
	return lastInsertId, nil
}
