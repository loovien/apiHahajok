/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/17 13:54
 * @description: 
 */

package criteria

type CommonCriteria struct {
	Condition string
	ConditionBind []interface{}
	Columns string
	Order string
	ReOrder bool
	OrCondition string
	OrConditionBind []interface{}
	Group string
	Having string
	HavingBind []interface{}
}
