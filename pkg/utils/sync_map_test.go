package utils

import (
	"encoding/json"
	ymal "gopkg.in/yaml.v3"
	"log"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"testing"
)

type Data struct {
	Name string
	Eq   string
	Pid  int
}

func TestHandleData(t *testing.T) {
	//config.ServerConfig.SetDeptFields("Name", "DID")
	//config.ServerConfig.SetDeptFields("Eq", "Order")
	//config.ServerConfig.SetDeptFields("Pid", "PlatformID")
	//data := Data{
	//	Name: "test",
	//	Eq:   "1",
	//	Pid:  1,
	//}
	//handleData := HandleDeptData(data)
	//fmt.Println(handleData)
	//utils2.JsonStringToStruct("[{\"third_name\":\"Name\",\"db_name\":\"DID\"},{\"third_name\":\"Eq\",\"db_name\":\"Order\"},{\"third_name\":\"Pid\",\"db_name\":\"PlatformID\"}]", &config.ServerConfig.UserFields)
	//fmt.Println(utils2.StructToJsonString(config.ServerConfig.UserFields))

	//test()
	testyaml()
}

type Config struct {
	SyncOptions struct {
		Sort int `yaml:"sort" json:"sort"`
	} `yaml:"syncOptions"`
	Fields []DeptField `yaml:"dept_fields" json:"dept_fields"`
}
type DeptField struct {
	ThirdName string `yaml:"third_name" json:"third_name,omitempty"`
	DBName    string `yaml:"db_name" json:"db_name,omitempty"`
}

func testyaml() {
	config := Config{SyncOptions: struct {
		Sort int `yaml:"sort" json:"sort"`
	}{Sort: 1},
		Fields: make([]DeptField, 0),
	}
	config.Fields = append(config.Fields, DeptField{
		ThirdName: "orgName",
		DBName:    "Name",
	})
	out, err := ymal.Marshal(&config)
	if err != nil {
		return
	}
	err = os.WriteFile("testConfig.yaml", out, 0644)
	if err != nil {
		log.Fatalf("写入文件失败: %v", err)
	}
}
func test() {
	m := make(map[string]interface{})
	m["name"] = "test"
	m["eq"] = 1
	jsonStr, err := json.Marshal(m)
	if err != nil {
		return
	}
	println(string(jsonStr))
	values := url.Values{}
	for k, v := range m {
		switch reflect.TypeOf(v).String() {
		case "string":
			values.Add(k, v.(string))
		case "int":
			values.Add(k, strconv.Itoa(v.(int)))
		}
	}
}
