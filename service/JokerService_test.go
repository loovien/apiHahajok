/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/16 20:05
 * @description: 
 */

package service

import (
	"testing"
	"github.com/vvotm/apiHahajok/models/request"
)

func TestGetLatestJokersList(tt *testing.T) {
	GetLatestJokersList(&request.ReqPage{1, 10})
}

func TestGetClassificationList(t *testing.T) {
	pageInfo := request.NewReqPage()
	resp, err := GetClassificationList(pageInfo)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}
