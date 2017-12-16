/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/16 10:19
 * @description: 
 */

package errhandle

type AbstractError struct {
	Code int
	Message string
}

func (p *AbstractError) GetCode() int  {
	return p.Code
}

func (p *AbstractError) Error() string {
	return p.Message
}

type PDOError struct {
	AbstractError
}
func NewPDOError(message string, code int) *PDOError {
	return &PDOError{
		AbstractError{code, message},
	}
}
