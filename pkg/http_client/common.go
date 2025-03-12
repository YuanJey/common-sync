package http_client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/YuanJey/go-log/pkg/log"
	"github.com/YuanJey/goutils2/pkg/utils"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func (c *HttpClient) Do(operationID, method, apiUrl, contentType string, req map[string]interface{}, header map[string]string) ([]byte, error) {
	body, err := handleBody(operationID, contentType, req)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(method, apiUrl, body)
	if err != nil {
		return nil, err
	}
	if contentType != "" {
		request.Header.Set("Content-Type", contentType)
	}
	for k, v := range header {
		request.Header.Set(k, v)
	}
	client := http.Client{Timeout: 5 * time.Second}
	httpResponse, err := client.Do(request)
	if err != nil {
		log.Error(operationID, "http client do failed", err.Error(), apiUrl)
		return nil, err
	}
	result, err := io.ReadAll(httpResponse.Body)
	if httpResponse.StatusCode != 200 {
		log.Error(operationID, "http client status code failed", apiUrl, string(result))
		return nil, utils.Wrap(errors.New(httpResponse.Status), "status code failed "+apiUrl+string(result))
	}
	return result, nil
}
func (c *HttpClient) Common(operationID, method, apiUrl, contentType string, req map[string]interface{}, resp interface{}, header map[string]string) error {
	body, err := handleBody(operationID, contentType, req)
	if err != nil {
		return err
	}
	request, err := http.NewRequest(method, apiUrl, body)
	if err != nil {
		return err
	}
	if contentType != "" {
		request.Header.Set("Content-Type", contentType)
	}
	for k, v := range header {
		request.Header.Set(k, v)
	}
	client := http.Client{Timeout: 5 * time.Second}
	httpResponse, err := client.Do(request)
	if err != nil {
		log.Error(operationID, "http client do failed", err.Error(), apiUrl)
		return err
	}
	result, err := io.ReadAll(httpResponse.Body)
	if httpResponse.StatusCode != 200 {
		log.Error(operationID, "http client status code failed", apiUrl, string(result))
		return utils.Wrap(errors.New(httpResponse.Status), "status code failed "+apiUrl+string(result))
	}
	err = utils.JsonStringToStruct(string(result), &resp)
	if err != nil {
		return err
	}
	return nil
}
func handleBody(operationID, contentType string, req map[string]interface{}) (io.Reader, error) {
	switch contentType {
	case ContentType_json:
		if req != nil {
			jsonStr, err := json.Marshal(req)
			if err != nil {
				log.Error(operationID, "json序列化失败:", err.Error())
				return nil, err
			}
			return strings.NewReader(string(jsonStr)), nil
		}
		return nil, nil
	case ContentType_form_urlencoded:
		values := url.Values{}
		for k, v := range req {
			switch reflect.TypeOf(v).String() {
			case "string":
				values.Add(k, v.(string))
			case "int":
				values.Add(k, strconv.Itoa(v.(int)))
			}
		}
		return strings.NewReader(values.Encode()), nil
	case ContentType_form_data:
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		for k, v := range req {
			err := writer.WriteField(k, v.(string))
			if err != nil {
				log.Error(operationID, "添加字段失败:", err.Error())
				return nil, err
			}
		}
		err := writer.Close()
		if err != nil {
			log.Error(operationID, "关闭失败:", err.Error())
			return nil, err
		}
		return body, nil
	}
	return nil, errors.New(fmt.Sprintf("不支持的contentType:%s", contentType))
}
