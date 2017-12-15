/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/16 2:28
 * @description: 
 */

package db

import (
	"testing"
	"time"
)

func TestGetConn(t *testing.T) {
	db := GetConn()
	result, err := db.Exec("insert into user (openId, createdAt) value(?, ?)", "罗文", time.Now().Unix())
	if err != nil {
		t.Log(err)
	}
	t.Log(result)
}
