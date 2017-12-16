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
	"github.com/labstack/gommon/log"
)

func GetLatestJokersList(pageInfo *request.ReqPage) (respJokerList response.RespJokerList, err error) {
	jokerDao := dao.NewJoker()
	conditions := fmt.Sprintf("1 = 1 limit %d, %d", (pageInfo.Page - 1) * pageInfo.Size, pageInfo.Size )
	jokerList, err := jokerDao.GetJokerList("*", conditions)
	log.Info(jokerList)
	if err != nil {
		return respJokerList, err
	}
	for _, joker := range jokerList{
		respJokerList.List = append(respJokerList.List, response.RespJokers{joker})
	}
	return respJokerList, nil
}
