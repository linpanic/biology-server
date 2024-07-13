package utils

import (
	"bytes"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var (
	httpTimeout = 10 * time.Second
	httpClient  = http.Client{
		Timeout: httpTimeout,
	}
)

type HttpResp struct {
	Err        error
	Data       []byte
	Localtion  string
	RequestUrl string
}

func HttpPostJson(api string, param interface{}) HttpResp {
	var result HttpResp
	var dataBytes []byte
	if param != nil {
		dataBytes, _ = json.Marshal(param)
	}
	reader := bytes.NewReader(dataBytes)
	request, err := http.NewRequest("POST", api, reader)
	if err != nil {
		log.Infof("请求:%s 出错:%s", api, err.Error())
		result.Err = err
		return result
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	resp, err := httpClient.Do(request)
	if err != nil {
		log.Infof("请求:%s 出错:%s", api, err.Error())
		result.Err = err
		return result
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	respBytes, _ := io.ReadAll(resp.Body)
	result.Data = respBytes
	return result
}

func HttpPostJsonAndHeader(api string, header map[string]string, param interface{}) HttpResp {
	var result HttpResp
	var dataBytes []byte
	if param != nil {
		dataBytes, _ = json.Marshal(param)
	}
	reader := bytes.NewReader(dataBytes)
	request, err := http.NewRequest("POST", api, reader)
	if err != nil {
		log.Infof("请求:%s 出错:%s", api, err.Error())
		result.Err = err
		return result
	}
	if len(header) > 0 {
		for k, v := range header {
			request.Header.Set(k, v)
		}
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	resp, err := httpClient.Do(request)
	if err != nil {
		log.Infof("请求:%s 出错:%s", api, err.Error())
		result.Err = err
		return result
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	respBytes, _ := io.ReadAll(resp.Body)
	result.Data = respBytes
	return result
}

func HttpGet(api string) HttpResp {
	var result HttpResp
	request, err := http.NewRequest("GET", api, nil)
	if err != nil {
		log.Infof("请求:%s 出错:%s", api, err.Error())
		result.Err = err
		return result
	}
	q := request.URL.Query()
	request.URL.RawQuery = q.Encode()
	resp, err := httpClient.Do(request)
	if err != nil {
		log.Infof("请求:%s 出错:%s", api, err.Error())
		result.Err = err
		return result
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	body, _ := io.ReadAll(resp.Body)
	result.Data = body
	result.RequestUrl = resp.Request.URL.String()
	return result
}

func HttpGetWithHeader(api string, h map[string]string) HttpResp {
	var result HttpResp
	request, err := http.NewRequest("GET", api, nil)
	if err != nil {
		log.Infof("请求:%s 出错:%s", api, err.Error())
		result.Err = err
		return result
	}

	for k, v := range h {
		request.Header.Set(k, v)
	}
	q := request.URL.Query()
	request.URL.RawQuery = q.Encode()
	resp, err := httpClient.Do(request)
	if err != nil {
		log.Infof("请求:%s 出错:%s", api, err.Error())
		result.Err = err
		return result
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	body, _ := io.ReadAll(resp.Body)
	result.Data = body
	result.RequestUrl = resp.Request.URL.String()
	return result
}

func HttpPostForm(api string, data url.Values) HttpResp {
	var result HttpResp
	r, err := http.NewRequest("POST", api, strings.NewReader(data.Encode())) // URL-encoded payload
	if err != nil {
		log.Infof("请求:%s 出错:%s", api, err.Error())
		result.Err = err
		return result
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := httpClient.Do(r)
	if err != nil {
		log.Infof("请求:%s 出错:%s", api, err.Error())
		result.Err = err
		return result
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	respBytes, _ := io.ReadAll(resp.Body)
	result.Data = respBytes
	return result
}
