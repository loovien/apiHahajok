/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/16 10:31
 * @description: 
 */

package errhandle

import "testing"

func TestNewPDOError(t *testing.T) {
	pdoErr := NewPDOError("数据库错误", ERROR_CODE)
	t.Log(pdoErr.Error())
}