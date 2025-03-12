package third_data

import (
	"common-sync/pkg/http_client"
	"errors"
	"fmt"
	"github.com/avast/retry-go/v4"
)

type GetThirdDataForApi interface {
	GetAllDept(operationID string) ([]ThirdDept, error)
	GetDeptByPage(operationID string, page int, pageSize int) ([]ThirdDept, error)
	GetAllUser(operationID string) ([]ThirdUser, error)
	GetUserByPage(operationID string, page int, pageSize int) ([]ThirdUser, error)
}
type APIConfig struct {
	Adders map[int]AddrInfo
}

func (a *APIConfig) AddAddrInfo(addr AddrInfo) {
	a.Adders[addr.Type] = addr
}

type CommonThirdData struct {
	httpClient *http_client.HttpClient
	APIConfig  APIConfig
	auth       *APIAuth
}
type GetAllDeptResp struct {
	result []ThirdDept
}

func (c *CommonThirdData) GetAllDept(operationID string) ([]ThirdDept, error) {
	result := GetAllDeptResp{}
	//var data []byte
	err := retry.Do(func() error {
		if addr, ok := c.APIConfig.Adders[http_client.Api_Data_All]; ok {
			if addr.Auth {
				c.auth.SetToken(operationID)
				addr.SetAuthToken(c.auth.Token)
			}
			err := c.httpClient.Common(operationID, addr.Method, addr.Url, addr.ContentType, addr.Req, &result, addr.Headers)
			if err != nil {
				return err
			}
			return nil
		}
		return errors.New("not found api APIConfig")
	}, retry.Attempts(3))
	return result.result, err
}

func (c *CommonThirdData) GetDeptByPage(operationID string, page int, pageSize int) ([]ThirdDept, error) {
	if addr, ok := c.APIConfig.Adders[http_client.Api_Data_Page]; ok {
		var all []ThirdDept
		var i = 1
		for {
			var result []ThirdDept
			err := c.httpClient.Common(operationID, addr.Method, fmt.Sprintln(addr.Url, i), addr.ContentType, addr.Req, &result, addr.Headers)
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
	return nil, errors.New("not found api APIConfig")
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
		APIConfig:  APIConfig{Adders: make(map[int]AddrInfo)},
	}
}
