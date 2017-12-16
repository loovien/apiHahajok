/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/16 10:31
 * @description: 
 */

package errhandle

import (
	"testing"
	//"errors"
)

func TestNewPDOError(t *testing.T) {
	pdoErr := NewPDOError("数据库错误", ERROR_CODE)
	//pdoErr := errors.New("hahah")
	defer func() {
		if r := recover(); r != nil {
			t.Log("not implement")
		}
	}()
	var absErr AbsErrer
	absErr = pdoErr.(AbsErrer)
	t.Log(absErr.GetCode())
}