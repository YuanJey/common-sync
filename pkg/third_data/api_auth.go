package third_data

import (
	"common-sync/pkg/http_client"
	"github.com/YuanJey/go-log/pkg/log"
)

type APIAuth struct {
	Url        string
	Method     string
	Req        map[string]interface{}
	Token      string
	httpClient *http_client.HttpClient
}
type APIAuthResp struct {
	Token string
}

func (a *APIAuth) SetToken(operationID string) {
	resp := APIAuthResp{}
	err := a.httpClient.Common(operationID, a.Method, a.Url, "", a.Req, &resp, nil)
	if err != nil {
		log.Error(operationID, "set token failed", err.Error(), *a)
		return
	}
	log.Info(operationID, "set token success", *a)
	a.Token = resp.Token
}
