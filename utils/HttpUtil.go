/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/16 0:15
 * @description: 
 */

package utils

import (
	"net/http"
	"io/ioutil"
)

// GET
func Get(URL string) ([]byte, error)  {
	resp, err := http.Get(URL)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}

