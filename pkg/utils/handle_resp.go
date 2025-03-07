package utils

import (
	"common-sync/pkg/config"
	"encoding/json"
	"fmt"
	"pl.ghgame.cn/gitea/yuanjie/db-sync-plugin/pkg/mysql/tables"
	"reflect"
	"strconv"
)

// {"data":[{"name":"aaa","age":18},{"name":"bbb","age":19}]}
func HandleOrgData(data []byte) []tables.TbLasDepartment {
	var departments []tables.TbLasDepartment
	var orgs interface{}
	err := json.Unmarshal([]byte(data), &orgs)
	if err != nil {
		panic(err)
	}
	dataValue := reflect.ValueOf(orgs)
	keys := dataValue.MapKeys()
	fmt.Println(keys)
	for _, key := range keys {
		value := dataValue.MapIndex(key)
		if "[]interface {}" == reflect.TypeOf(value.Interface()).String() {
			for _, org := range value.Interface().([]interface{}) {
				department := tables.TbLasDepartment{}
				fmt.Println(reflect.ValueOf(org))
				orgKeys := reflect.ValueOf(org).MapKeys()
				for _, orgKey := range orgKeys {
					k := reflect.ValueOf(orgKey)
					v := reflect.ValueOf(org).MapIndex(key)
					fmt.Println(k)
					fmt.Println(k.Kind().String())
					for i := range config.ServerConfig.DeptFields {
						fmt.Println()
						if k.String() == config.ServerConfig.DeptFields[i].ThirdName {
							reflect.ValueOf(&department).Elem().FieldByName(config.ServerConfig.DeptFields[i].DBName).Set(reflect.ValueOf(v))
						}
					}
				}
				departments = append(departments, department)
			}
		}
	}
	//for _, org := range orgs.([]interface{}) {
	//	department := handleDeptData(org)
	//	departments = append(departments, department)
	//}
	return departments
}
func handleDeptData(data interface{}) tables.TbLasDepartment {
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
