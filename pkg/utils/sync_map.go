package utils

import (
	"common-sync/pkg/config"
	"pl.ghgame.cn/gitea/yuanjie/db-sync-plugin/pkg/mysql/tables"
	"reflect"
	"strconv"
)

func HandleDeptData(data interface{}) tables.TbLasDepartment {
	department := tables.TbLasDepartment{}
	value := reflect.ValueOf(data)
	for i := range config.ServerConfig.DeptFields {
		thirdValue := value.FieldByName(config.ServerConfig.DeptFields[i].ThirdName).Interface()
		thirdT := reflect.TypeOf(thirdValue).String()
		dbT := reflect.TypeOf(reflect.ValueOf(&department).Elem().FieldByName(config.ServerConfig.DeptFields[i].DBName).Interface()).String()
		if dbT == thirdT {
			reflect.ValueOf(&department).Elem().FieldByName(config.ServerConfig.DeptFields[i].DBName).Set(reflect.ValueOf(thirdValue))
		} else {
			if thirdT == "string" && dbT == "int" {
				thirdValueInt, _ := strconv.Atoi(thirdValue.(string))
				reflect.ValueOf(&department).Elem().FieldByName(config.ServerConfig.DeptFields[i].DBName).Set(reflect.ValueOf(thirdValueInt))
			}
			if thirdT == "int" && dbT == "string" {
				thirdValueStr := strconv.Itoa(thirdValue.(int))
				reflect.ValueOf(&department).Elem().FieldByName(config.ServerConfig.DeptFields[i].DBName).Set(reflect.ValueOf(thirdValueStr))
			}
		}
	}
	return department
}
func HandleUserData(data interface{}) (tables.TbLasUser, tables.TbLasDepartmentUser) {
	lasUser := tables.TbLasUser{}
	lasDeptUser := tables.TbLasDepartmentUser{}
	value := reflect.ValueOf(data)
	for i := range config.ServerConfig.DeptFields {
		thirdValue := value.FieldByName(config.ServerConfig.DeptFields[i].ThirdName).Interface()
		thirdT := reflect.TypeOf(thirdValue).String()
		dbT := reflect.TypeOf(reflect.ValueOf(&lasUser).Elem().FieldByName(config.ServerConfig.DeptFields[i].DBName).Interface()).String()
		if dbT == thirdT {
			reflect.ValueOf(&lasUser).Elem().FieldByName(config.ServerConfig.DeptFields[i].DBName).Set(reflect.ValueOf(thirdValue))
			reflect.ValueOf(&lasDeptUser).Elem().FieldByName(config.ServerConfig.DeptFields[i].DBName).Set(reflect.ValueOf(thirdValue))
		} else {
			if thirdT == "string" && dbT == "int" {
				thirdValueInt, _ := strconv.Atoi(thirdValue.(string))
				reflect.ValueOf(&lasUser).Elem().FieldByName(config.ServerConfig.DeptFields[i].DBName).Set(reflect.ValueOf(thirdValueInt))
				reflect.ValueOf(&lasDeptUser).Elem().FieldByName(config.ServerConfig.DeptFields[i].DBName).Set(reflect.ValueOf(thirdValueInt))
			}
			if thirdT == "int" && dbT == "string" {
				thirdValueStr := strconv.Itoa(thirdValue.(int))
				reflect.ValueOf(&lasUser).Elem().FieldByName(config.ServerConfig.DeptFields[i].DBName).Set(reflect.ValueOf(thirdValueStr))
				reflect.ValueOf(&lasDeptUser).Elem().FieldByName(config.ServerConfig.DeptFields[i].DBName).Set(reflect.ValueOf(thirdValueStr))
			}
		}
	}
	return lasUser, lasDeptUser
}
