/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/26 21:47
 * @description: 
 */

package validate

import (
	"github.com/labstack/echo"
	"github.com/vvotm/apiHahajok/models/request"
	"errors"
)

type RepliesValidate struct {
	validate *echo.Validator
}

func NewRepliesValidate() *RepliesValidate  {
	return &RepliesValidate{}
}

func (rv *RepliesValidate) Validate(reqReplies *request.ReqReplies) error {
	if reqReplies.Uid <= 0 {
		return errors.New("UID 不能为空")
	}
	if reqReplies.Content == "" {
		return errors.New("Content 不能为空")
	}
	if reqReplies.JokerId <= 0 {
		return errors.New("jokerId 不能为空")
	}
	return nil
}
