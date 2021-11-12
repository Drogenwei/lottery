package helpers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

/**
 * 获取 url 内容
 */
func Get(url string) string {
	if url == "" {
		panic("获取链接为空")
	}

	resp, err := http.Get(url)
	if err != nil {
		panic("获取链接内容错误" + err.Error())
	}
	defer resp.Body.Close()

	// 去读数据内容为 bytes
	dataBytes, err := ioutil.ReadAll(resp.Body)
	if err !=nil{
		panic("ioutil.ReadAll error : "+err.Error())
	}
	// 转码
	utf8 := ConvertToString(string(dataBytes), "gbk", "utf-8")

	return utf8
}

// 发送post请求
func SendPost(url string, jsonStr []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
fmt.Println(string(body))
	return body, nil
}

// 发送get请求
func SendGet(url string, header map[string]string) []byte {
	client := &http.Client{}
	req,_ := http.NewRequest("GET",url,nil)
	for k, v := range header {
		req.Header.Add(k, v)
	}
	resp,_ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)

	return body
}