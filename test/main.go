package main

import (
	"common-sync/pkg/config"
	"common-sync/pkg/utils"
	"encoding/json"
	"fmt"
	"github.com/YuanJey/commonHttpClient/pkg/params"
	ymal "gopkg.in/yaml.v3"
	"os"
	"reflect"
)

func main() {
	//save()
	//resp()
	config.ServerConfig.DeptFields = []config.DeptField{
		{
			DBName:    "Name",
			ThirdName: "name",
		},
		{
			DBName:    "Order",
			ThirdName: "age",
		},
	}
	str := "{\"data\":[{\"name\":\"aaa\",\"age\":18},{\"name\":\"bbb\",\"age\":19}]}"
	data := utils.HandleOrgData([]byte(str))
	fmt.Println(data)
}

func resp() {
	str := "{\"data\":[{\"name\":\"aaa\",\"age\":18},{\"name\":\"bbb\",\"age\":19}]}"
	//str := "{\"name\":\"aaa\",\"age\":18}"
	var user interface{}
	err := json.Unmarshal([]byte(str), &user)
	if err != nil {
		panic(err)
	}
	reflectStruct(user)
}
func reflectStruct(val interface{}) {
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.Map:
		for _, key := range v.MapKeys() {
			val := v.MapIndex(key)
			fmt.Printf("Key: %s, Value: %v\n", key.String(), val.Interface())
			reflectStruct(val.Interface())
		}
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			val := v.Index(i)
			fmt.Printf("Index: %d, Value: %v\n", i, val.Interface())
			reflectStruct(val.Interface())
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			field := v.Type().Field(i)
			val := v.Field(i)
			fmt.Printf("Field: %s, Value: %v\n", field.Name, val.Interface())
			reflectStruct(val.Interface())
		}
	default:
		fmt.Printf("Value: %v\n", v.Interface())
	}
}
func save() {
	h := make(map[string]string)
	h["Content-Type"] = "application/json"
	m := make(map[string]interface{})
	m["size"] = 1
	m["page"] = 1
	requestConfig := params.RequestConfig{
		Method:      "GET",
		BodyType:    1,
		Url:         "https://www.baidu.com",
		ContentType: "application/json",
		Headers:     h,
		Req:         m,
		PageConf: params.PageConfig{
			IsPage:    true,
			Page:      1,
			PageField: "page",
			PageSize:  100,
			SizeField: "size",
		},
	}
	var l []params.RequestConfig
	l = append(l, requestConfig)
	out, err := ymal.Marshal(&l)
	if err != nil {
		return
	}
	err = os.WriteFile("apiConfig.yaml", out, 0644)
	if err != nil {
		return
	}
}
