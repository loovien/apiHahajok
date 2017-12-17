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
		CommonCriteria: criteria.CommonCriteria{
			Condition: "status = ?",
			ConditionBind: []interface{}{1},
			Group: "status",
		},
		Page: 1,
		Size: 10,
	}
	list, err := joker.GetJokerList(criteria)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(list)
}

func TestJoker_Count(t *testing.T) {
	joker := NewJoker()
	criteria := criteria.CommonCriteria{
	}
	cnt := joker.Count(criteria)
	t.Log(cnt)
}

func TestJoker_DeleteJoker(t *testing.T) {
	joker := NewJoker()
	t.Log(joker.DeleteJoker("id = 1"))
}
