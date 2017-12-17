/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/16 18:48
 * @description: 
 */

package dao

import (
	"testing"
	"github.com/vvotm/apiHahajok/dao/criteria"
)

func TestJoker_GetJokerList(t *testing.T) {
	joker := NewJoker()
	criteria := criteria.PageCriteria{
		Condition: "status = ?",
		ConditionBind: []interface{}{1},
		Page: 1,
		Size: 10,
		Group: "status",
	}
	list, err := joker.GetJokerList(criteria)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(list)
}

func TestJoker_Count(t *testing.T) {
	joker := NewJoker()
	cnt := joker.Count("1=1")
	t.Log(cnt)
}
