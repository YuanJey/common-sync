package third_data

import (
	"common-sync/pkg/http_client"
	"errors"
	"fmt"
)

type GetThirdDataForApi interface {
	GetAllDept(operationID string) ([]ThirdDept, error)
	GetDeptByPage(operationID string, page int, pageSize int) ([]ThirdDept, error)
	GetAllUser(operationID string) ([]ThirdUser, error)
	GetUserByPage(operationID string, page int, pageSize int) ([]ThirdUser, error)
}
type Config struct {
	adders map[int]addrInfo
}
type addrInfo struct {
	Url         string
	Method      string
	contentType string
	Headers     map[string]string
	Type        int
	Req         map[string]interface{}
}
type CommonThirdData struct {
	httpClient *http_client.HttpClient
	config     Config
}

func (c *CommonThirdData) GetAllDept(operationID string) ([]ThirdDept, error) {
	if addr, ok := c.config.adders[http_client.Api_Data_All]; ok {
		var result []ThirdDept
		err := c.httpClient.Common(operationID, addr.Method, addr.Url, addr.contentType, addr.Req, &result, addr.Headers)
		if err != nil {
			return nil, err
		}
		return result, nil
	}
	return nil, errors.New("not found api config")
}

func (c *CommonThirdData) GetDeptByPage(operationID string, page int, pageSize int) ([]ThirdDept, error) {
	if addr, ok := c.config.adders[http_client.Api_Data_Page]; ok {
		var all []ThirdDept
		var i = 1
		for {
			var result []ThirdDept
			err := c.httpClient.Common(operationID, addr.Method, fmt.Sprintln(addr.Url, i), addr.contentType, addr.Req, &result, addr.Headers)
			if err != nil {
				return nil, err
			}
			i++
			if len(result) > 0 {
				all = append(all, result...)
			}
			if len(result) < 500 {
				break
			}
		}
		return all, nil
	}
	return nil, errors.New("not found api config")
}

func (c *CommonThirdData) GetAllUser(operationID string) ([]ThirdUser, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CommonThirdData) GetUserByPage(operationID string, page int, pageSize int) ([]ThirdUser, error) {
	//TODO implement me
	panic("implement me")
}

func NewCommonThirdData() *CommonThirdData {
	return &CommonThirdData{
		httpClient: &http_client.HttpClient{},
		config:     Config{adders: make(map[int]addrInfo)},
	}
}
