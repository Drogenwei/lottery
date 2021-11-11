package helpers

import (
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