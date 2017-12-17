/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/16 2:28
 * @description: 
 */

package db

import (
	"testing"
)

func TestGetConn(t *testing.T) {
	db := GetConn()
	user := struct {
		Id int
		OpenId string
		UnionId string
		Nickname string
	}{}
	db.Raw("select * from user where id = 1").Scan(&user)
	t.Log(user)
}
