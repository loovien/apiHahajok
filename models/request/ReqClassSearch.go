/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/17 22:15
 * @description: 
 */

package request

type ReqClassSearch struct {
	ReqPage
	Name string `json:"name" form:"name" query:"name"`
}

func NewReqClassSearch() *ReqClassSearch {
	return &ReqClassSearch{
		ReqPage: *NewReqPage(),
		Name: "",
	}
}
