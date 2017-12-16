/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/16 18:48
 * @description: 
 */

package dao

import (
	"testing"
)

func TestJoker_GetJokerList(t *testing.T) {
	joker := NewJoker()
	list, err := joker.GetJokerList("id, classId, title, imageList", "")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(list)
}

