/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/16 0:01
 * @description:  wechat Util
 */

package utils

import (
	"fmt"
	"github.com/vvotm/apiHahajok/constants"
	"encoding/json"
)

// AUTH_URL
const AUTH_URL  = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"

// GetOpenID
func GetOpenID(jsCode string) (map[string]interface{}, error) {
	authURL := fmt.Sprintf(AUTH_URL, constants.APP_ID, constants.APP_SECRET, jsCode)
	bytes, err := Get(authURL)
	if err != nil {
		return nil, err
	}
	resp := make(map[string]interface{})
	json.Unmarshal(bytes, &resp)
	return resp, nil
}
