package config

import (
	"github.com/YuanJey/commonHttpClient/pkg/params"
	"github.com/YuanJey/go-log/pkg/log"
	"github.com/YuanJey/goconf/pkg/config"
	ymal "gopkg.in/yaml.v3"
	"os"
)

var ServerConfig Config

const (
	ApiConfigKey_Dept = "deptData"
	ApiConfigKey_User = "userData"
)

type Config struct {
	WPS struct {
		Addr       string `yaml:"addr"`
		CompanyId  string `yaml:"companyId"`
		PlatformId string `yaml:"platformId"`
	} `yaml:"wps"`
	Api struct {
		AdminKey string `yaml:"admin_key" env:"API_ADMIN_KEY"`
		PORT     string `yaml:"port" env:"HTTP_KPORT"`
	}
	DeptFields  []DeptField `yaml:"dept_fields" json:"dept_fields"`
	UserFields  []UserField `yaml:"user_fields" json:"user_fields"`
	SyncOptions struct {
		Sort int `yaml:"sort" json:"sort"`
	} `yaml:"syncOptions"`
	ApiConfig map[string]params.RequestConfig `yaml:"apiConfig" json:"apiConfig"`
}

func (c *Config) Save(operationID, fileName string) error {
	out, err := ymal.Marshal(c)
	if err != nil {
		log.Error(operationID, "Marshal err ", err.Error())
		return err
	}
	err = os.WriteFile(fileName, out, 0644)
	if err != nil {
		log.Error(operationID, "Save config file error: ", err.Error())
		return err
	}
	return nil
}
func (c *Config) SetDeptFields(thirdField string, dbField string) {
	for i := range c.DeptFields {
		if c.DeptFields[i].DBName == dbField {
			c.DeptFields[i].ThirdName = thirdField
			return
		}
	}
	c.DeptFields = append(c.DeptFields, DeptField{
		ThirdName: thirdField,
		DBName:    dbField,
	})
}
func (c *Config) SetUserFields(thirdField string, dbField string) {
	for i := range c.UserFields {
		if c.UserFields[i].DBName == dbField {
			c.UserFields[i].ThirdName = thirdField
			return
		}
	}
	c.UserFields = append(c.UserFields, UserField{
		ThirdName: thirdField,
		DBName:    dbField,
	})
}

type DeptField struct {
	ThirdName string `yaml:"third_name" json:"third_name,omitempty"`
	DBName    string `yaml:"db_name" json:"db_name,omitempty"`
}
type UserField struct {
	ThirdName string `yaml:"third_name" json:"third_name,omitempty"`
	DBName    string `yaml:"db_name" json:"db_name,omitempty"`
}

func init() {
	config.UnmarshalConfig(&ServerConfig, "config.yaml")
}
