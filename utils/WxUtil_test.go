/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/16 0:47
 * @description: 
 */

package utils

import (
	"testing"
	"net/http"
	"io/ioutil"
)

func TestGetOpenID(t *testing.T) {
	var url string = "https://api.weixin.qq.com/sns/jscode2session?appid=wxe32a496a4c53509d&secret=300618335bf4a1442b9ae525db9980ac&js_code=JSCODE&grant_type=authorization_code"
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	t.Log(string(content))
}
