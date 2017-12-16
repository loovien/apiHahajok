/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/16 18:02
 * @description: 
 */

package dao

type Joker struct {
	Id int `json:"id"`
	Uid int `json:"uid"`
	ClassId int `json:"classId"`
	Title string `json:"title"`
	Content string `json:"content"`
	ImageList []string `json:"imageList"`
	MediaUrl string `json:"mediaUrl"`
	Replies int `json:"replies"`
	Status int `json:"status"`
	CreatedAt int `json:"createdAt"`
	PassedAt int `json:"passedAt"`
	UpdatedAt int `json:"updatedAt"`
}
