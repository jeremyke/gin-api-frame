package http_request

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/imroc/req"
)

// HttpRequest http 请求
func HttpRequest(requestType string, url string, reqData map[string]interface{}) (map[string]interface{}, error) {
	client := &http.Client{}
	bytesData, _ := json.Marshal(reqData)
	req, _ := http.NewRequest(requestType, url, bytes.NewReader(bytesData))
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)

	resData := map[string]interface{}{}
	err := json.Unmarshal(body, &resData)
	if err != nil {
		return nil, err
	}
	return resData, nil
}

// HttpPost http 请求
func HttpPost(url string, data req.Param) (res []byte, err error) {
	header := req.Header{
		"Accept": "application/json",
	}
	r, err := req.Post(url, header, req.BodyJSON(data))
	if err != nil {
		return
	}

	resBody, err := r.ToBytes()
	if err != nil {
		return
	}
	defer r.Response().Body.Close()
	return resBody, err
}
