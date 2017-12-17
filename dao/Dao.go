/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/17 13:36
 * @description: 
 */

package dao

import (
	"github.com/vvotm/apiHahajok/dao/criteria"
	"github.com/jinzhu/gorm"
)

func ApplyPageQuery(orm *gorm.DB, p criteria.PageCriteria) *gorm.DB {
	orm = ApplyCommonQuery(orm, p.CommonCriteria)
	if p.Page > 0 {
		offset := (p.Page - 1) * p.Size
		orm = orm.Offset(offset)
	}
	orm = orm.Limit(p.Size)
	return orm
}

func ApplyCommonQuery(orm *gorm.DB, p criteria.CommonCriteria) *gorm.DB {
	if p.Condition != "" {
		orm = orm.Where(p.Condition, p.ConditionBind...)
	}
	if p.Columns != "" {
		orm = orm.Select(p.Columns)
	}
	if p.Order != "" {
		orm = orm.Order(p.Order, p.ReOrder)
	}
	if p.Group != "" {
		orm = orm.Group(p.Group)
	}
	if p.Having != "" {
		orm = orm.Having(p.Having, p.HavingBind...)
	}
	return orm
}
