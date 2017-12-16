/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/16 18:03
 * @description: 
 */

package dao

type Classification struct {
	Id  int `json:"id"`
	Name  string `json:"name"`
	Icon string `json:"icon"`
	CreatedAt int `json:"createdAt"`
	UpdatedAt int `json:"updatedAt"`
}
