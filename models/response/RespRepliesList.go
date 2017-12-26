/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/26 20:20
 * @description: 
 */

package response

type RespRepliesList struct {
	RespPage
	List []RespReplies `json:"list"`
}
