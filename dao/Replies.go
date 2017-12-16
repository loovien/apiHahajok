/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/16 18:02
 * @description: 
 */

package dao

type Replies struct {
	Id int `json:"id"`
	JokerId int `json:"jokerId"`
	Uid int `json:"uid"`
	Content string `json:"content"`
	ImageList []string `json:"imageList"`
	Status int `json:"status"`
	CreatedAt int `json:"createdAt"`
	PassedAt int `json:"passedAt"`
	UpdatedAt int `json:"updatedAt"`
}
