/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/16 18:08
 * @description: 
 */

package response

type RespJokerList struct {
	RespPage
	List []RespJokers `json:"list"`
}
