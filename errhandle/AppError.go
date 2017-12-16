/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/16 10:19
 * @description: 
 */

package errhandle

type AbstractError struct {
	Message string
}
func (p *AbstractError) Error() string {
	return p.Message
}

type PDOError struct {
	AbstractError
}
func NewPDOError(message string) *PDOError {
	return &PDOError{
		AbstractError{message},
	}
}
