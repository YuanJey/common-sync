package http_client

import (
	"encoding/json"
	"errors"
	"github.com/YuanJey/go-log/pkg/log"
	"github.com/YuanJey/goutils2/pkg/utils"
	"io"
	"net/http"
	"strings"
	"time"
)

type HttpClient struct{}

func (c *HttpClient) Post(operationID, url string, req interface{}, resp interface{}, sign Sign) error {
	body := strings.NewReader("")
	if req != nil {
		jsonStr, err := json.Marshal(req)
		if err != nil {
			return err
		}
		body = strings.NewReader(string(jsonStr))
	}
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	if sign != nil {
		sign.Sign(request)
	}
	client := http.Client{Timeout: 5 * time.Second}
	httpResponse, err := client.Do(request)
	if err != nil {
		log.Error(operationID, "http client do failed", err.Error(), url)
		return err
	}
	result, err := io.ReadAll(httpResponse.Body)
	if httpResponse.StatusCode != 200 {
		log.Error(operationID, "http client status code failed", url, string(result))
		return utils.Wrap(errors.New(httpResponse.Status), "status code failed "+url+string(result))
	}
	err = utils.JsonStringToStruct(string(result), &resp)
	if err != nil {
		return err
	}
	return nil
}
func (c *HttpClient) Get(operationID, url string, resp interface{}, sign Sign) error {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	if sign != nil {
		sign.Sign(request)
	}
	client := http.Client{Timeout: 5 * time.Second}
	httpResponse, err := client.Do(request)
	if err != nil {
		log.Error(operationID, "http client do failed", err.Error(), url)
		return err
	}
	defer httpResponse.Body.Close()
	result, err := io.ReadAll(httpResponse.Body)
	if httpResponse.StatusCode != 200 {
		log.Error(operationID, "http client status code failed", url, string(result))
		return utils.Wrap(errors.New(httpResponse.Status), "status code failed "+url+string(result))
	}
	err = utils.JsonStringToStruct(string(result), &resp)
	if err != nil {
		return err
	}
	return nil
}
