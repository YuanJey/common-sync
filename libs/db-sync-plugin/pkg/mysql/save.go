package mysql

import (
	"github.com/YuanJey/goutils2/pkg/utils"
	"pl.ghgame.cn/gitea/yuanjie/db-sync-plugin/pkg/mysql/tables"
)

func BatchSaveTbLasDepartment(list []tables.TbLasDepartment) error {
	if list == nil {
		return nil
	}
	tableName := list[0].TableName()
	return utils.Wrap(Database.db.Table(tableName).Save(list).Error, "batch save failed "+tableName)
}
func BatchSaveTbLasUser(list []tables.TbLasUser) error {
	if list == nil {
		return nil
	}
	tableName := list[0].TableName()
	return utils.Wrap(Database.db.Table(tableName).Save(list).Error, "batch save failed "+tableName)
}
func BatchSaveTbLasDepartmentUser(list []tables.TbLasDepartmentUser) error {
	if list == nil {
		return nil
	}
	tableName := list[0].TableName()
	return utils.Wrap(Database.db.Table(tableName).Save(list).Error, "batch save failed "+tableName)
}
func SaveTbLasDepartment(list *tables.TbLasDepartment) error {
	if list == nil {
		return nil
	}
	tableName := list.TableName()
	return utils.Wrap(Database.db.Table(tableName).Save(list).Error, "save failed "+tableName)
}
func SaveTbLasUser(list *tables.TbLasUser) error {
	if list == nil {
		return nil
	}
	tableName := list.TableName()
	return utils.Wrap(Database.db.Table(tableName).Save(list).Error, "save failed "+tableName)
}
func SaveTbLasDepartmentUser(list *tables.TbLasDepartmentUser) error {
	if list == nil {
		return nil
	}
	tableName := list.TableName()
	return utils.Wrap(Database.db.Table(tableName).Save(list).Error, "save failed "+tableName)
}
