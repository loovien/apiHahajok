/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/16 13:45
 * @description: 
 */

package request

type ReqPage struct {
	Page int `json:"page" form:"page" query:"page"`
	Size int `json:"size" form:"size" query:"size"`
}

func NewReqPage() *ReqPage {
	return &ReqPage{
		Page: 1,
		Size: 10,
	}
}
