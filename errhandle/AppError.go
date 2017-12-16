/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/16 10:19
 * @description: 
 */

package errhandle

type AbsErrer interface {
	GetCode() int
	Error() string
	ErrorMsg() string
}

type SuperError struct {
	Code int `json:"code"`
	Message string `json:"message"`
	ErrMsg string `json:"errmsg"`

}

func (p *SuperError) GetCode() int  {
	return p.Code
}

func (p *SuperError) Error() string {
	return p.Message
}

func (p *SuperError) ErrorMsg() string {
	return p.ErrMsg
}

type PDOError struct {
	*SuperError
}
func NewPDOError(message string, code int, errmsg string) AbsErrer {
	return &PDOError{
		&SuperError{code, message, errmsg},
	}
}
