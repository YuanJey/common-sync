package mysql

import (
	"fmt"
	"github.com/YuanJey/goutils2/pkg/utils"
	"gorm.io/gorm/schema"
	"pl.ghgame.cn/gitea/yuanjie/db-sync-plugin/pkg/mysql/tables"
)

func BatchInsertTbLasDepartment(list []tables.TbLasDepartment) error {
	if list == nil {
		return nil
	}
	tableName := list[0].TableName()
	Database.db.Config.CreateBatchSize = 10
	return utils.Wrap(Database.db.Table(tableName).Create(list).Error, "batch insert failed "+tableName)
}

func BatchInsertTbLasDepartmentIncrement(list []tables.TbLasDepartmentIncrement) error {
	if list == nil {
		return nil
	}
	tableName := list[0].TableName()
	return utils.Wrap(Database.db.Table(tableName).Create(list).Error, "batch insert failed "+tableName)
}
func BatchInsertTbLasUser(list []tables.TbLasUser) error {
	if list == nil {
		return nil
	}
	tableName := list[0].TableName()
	return utils.Wrap(Database.db.Table(tableName).Create(list).Error, "batch insert failed "+tableName)
}
func BatchInsertTbLasDepartmentUser(list []tables.TbLasDepartmentUser) error {
	if list == nil {
		return nil
	}
	tableName := list[0].TableName()
	return utils.Wrap(Database.db.Table(tableName).Create(list).Error, "batch insert failed "+tableName)
}

func BatchInsertTbLasDepartmentUserIncrement(list []tables.TbLasDepartmentUserIncrement) error {
	if list == nil {
		return nil
	}
	tableName := list[0].TableName()
	return utils.Wrap(Database.db.Table(tableName).Create(list).Error, "batch insert failed "+tableName)
}

func BatchInsertTbLasUserIncrement(list []tables.TbLasUserIncrement) error {
	if list == nil {
		return nil
	}
	tableName := list[0].TableName()
	err := utils.Wrap(Database.db.Table(tableName).Create(list).Error, "batch insert failed "+tableName)
	if err != nil {
		fmt.Println("插入失败 ", err.Error())
		return err
	}
	return nil
}
func BatchInsert(list []schema.Tabler) error {
	tableName := list[0].TableName()
	return utils.Wrap(Database.db.Table(tableName).Create(list).Error, "batch insert failed "+tableName)
}
func Install(data schema.Tabler) error {
	return utils.Wrap(Database.db.Table(data.TableName()).Create(data).Error, "install failed "+data.TableName())
}

func Clean(data schema.Tabler) error {
	return utils.Wrap(Database.db.Table(data.TableName()).Delete(data).Error, "clean failed "+data.TableName())
}
