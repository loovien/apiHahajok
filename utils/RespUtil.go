/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/15 23:51
 * @description: app uitls
 */

package utils

// GetCommonResp
// build a common response json data
func GetCommonResp(data interface{}, code int, message string) (interface{}){
	return struct {
		Code int `json:"code"`
		Message string `json:"message"`
		Data interface{} `json:"data"`
	}{Code: code, Message: message, Data:data}
}

