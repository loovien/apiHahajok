/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/16 20:05
 * @description: 
 */

package service

import (
	"testing"
	"github.com/vvotm/apiHahajok/models/response"
	"github.com/vvotm/apiHahajok/dao"
	"github.com/vvotm/apiHahajok/models/request"
)

func TestGetLatestJokersList(tt *testing.T) {

	b := []response.RespJokers{}
	joker := dao.Joker{Id:1, Uid:23}
	respJ := response.RespJokers{DoJoker:joker}
	b = append(b, respJ)
	tt.Log(b)
	GetLatestJokersList(&request.ReqPage{1, 10})
}
