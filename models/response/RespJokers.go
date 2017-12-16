/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/16 13:57
 * @description: 
 */

package response

import "github.com/vvotm/apiHahajok/dao"

type RespJokers struct {
	dao.Joker
	Member RespUser `json:"member"`
	Classification RespClassification `json:"classification"`
}
