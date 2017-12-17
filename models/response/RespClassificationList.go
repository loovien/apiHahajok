/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/17 21:38
 * @description: 
 */

package response

type RespClassificationList struct {
	RespPage
	List []RespClassification `json:"list"`
}
