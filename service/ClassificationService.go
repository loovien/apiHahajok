/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/17 21:37
 * @description: 
 */

package service

import (
	"github.com/vvotm/apiHahajok/models/response"
	"github.com/vvotm/apiHahajok/models/request"
	"github.com/vvotm/apiHahajok/dao/criteria"
	"github.com/vvotm/apiHahajok/dao"
)

func GetClassificationList(pageInfo *request.ReqPage) (respClassificationList response.RespClassificationList, err error) {
	respClassificationList.RespPage.Size = pageInfo.Size
	 commonCriteria := criteria.CommonCriteria{
		Condition: "status = 1",
	}
	classificationDao := dao.NewClassification()
	respClassificationList.RespPage.Cnt = classificationDao.Count(commonCriteria)

	pageCriteria := criteria.PageCriteria{
		CommonCriteria: commonCriteria,
		Page: pageInfo.Page,
		Size: pageInfo.Size,
	}
	classificationList, err := classificationDao.GetClassificationList(pageCriteria)
	if err != nil {
		return respClassificationList, err
	}
	for _, classification := range classificationList {
		respClassificationList.List = append(respClassificationList.List, response.RespClassification{
			Classification: classification,
		})
	}
	return respClassificationList, nil
}
