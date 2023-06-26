package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// GetJSON 发送get请求 返回json并绑定response结构
func GetJSON(url string, response interface{}) error {

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(response)
}

// HttpPost 发送Post请求
func HttpPost(url string, body interface{}) (data []byte, err error) {
	bodyStr, err := json.Marshal(body)
	if err != nil {
		return
	}
	resp, err := http.Post(url, "application/json", bytes.NewReader(bodyStr))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	data, err = ioutil.ReadAll(resp.Body)
	return
}

// HttpGet 发送Get请求
func HttpGet(url string) (data []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	data, err = ioutil.ReadAll(resp.Body)
	return
}
