/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/17 13:29
 * @description: generate page query struct
 */

package criteria

type PageCriteria struct {
	Condition string
	ConditionBind []interface{}
	Columns string
	Order string
	ReOrder bool
	Page int
	Size int
	OrCondition string
	OrConditionBind []interface{}
	Group string
	Having string
	HavingBind []interface{}
}
