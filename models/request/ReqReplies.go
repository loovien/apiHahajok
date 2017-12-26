/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/26 20:52
 * @description: 
 */

package request

type ReqReplies struct {
	JokerId int `json:"jokerId" form:"jokerId" query:"jokerId"`
	Uid int `json:"uid" form:"uid" query:"uid"`
	Content string `json:"content" form:"content" query:"content"`
	ImageList string `json:"imageList" form:"imageList" query:"imageList"`
}

func NewReqReplies() *ReqReplies {
	return &ReqReplies{}
}
