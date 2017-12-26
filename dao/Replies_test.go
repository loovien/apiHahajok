/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/26 21:22
 * @description: 
 */

package dao

import (
	"testing"
	"time"
)

func TestReplies_Insert(t *testing.T) {
	repliesDao := NewReplies()
	nowTime := int(time.Now().Unix())
	repliesDao.ImageList = "haha阿斯利的肌肤三分"
	repliesDao.Content = "th问哦解放路看金利科技 is a good func"
	repliesDao.Status = 1
	repliesDao.Uid = 11
	repliesDao.JokerId = 1212
	repliesDao.CreatedAt = nowTime
	repliesDao.UpdatedAt = nowTime
	repliesDao.PassedAt = nowTime

	t.Log(repliesDao)
	t.Log(repliesDao.Insert())

}
