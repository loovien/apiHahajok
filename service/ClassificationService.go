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

func GetClassificationList(classSearch *request.ReqClassSearch) (respClassificationList response.RespClassificationList, err error) {
	respClassificationList.RespPage.Size = classSearch.Size
	conditions := "status = 1"
	if classSearch.Name != "" {
		conditions += " and name like '" + classSearch.Name + "%'"
	}
	 commonCriteria := criteria.CommonCriteria{
		Condition: conditions,
	}
	classificationDao := dao.NewClassification()
	respClassificationList.RespPage.Cnt = classificationDao.Count(commonCriteria)

	pageCriteria := criteria.PageCriteria{
		CommonCriteria: commonCriteria,
		Page: classSearch.Page,
		Size: classSearch.Size,
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
