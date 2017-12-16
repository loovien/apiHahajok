/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/16 11:40
 * @description: 
 */

package request

type ReqUserInfo struct {
	OpenId string `json:"openid" form:"openid" query:"openid"`
	UnionId string `json:"unionid" form:"unionid" query:"unionid"`
	Nickname string `json:"nickname" form:"nickname" query:"nickname"`
	Avatar string `json:"avatar" form:"avatar" query:"avatar"`
	Gender string `json:"gender" form:"gender" query:"gender"`
	Country string `json:"country" form:"country" query:"country"`
	Province string `json:"province" form:"province" query:"province"`
	City string `json:"city" form:"city" query:"city"`
	Lang string `json:"lang" form:"lang" query:"lang"`
}
