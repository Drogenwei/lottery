package helpers

import (
	"encoding/json"
	"github.com/axgle/mahonia"
	"math/rand"
	"time"
)

/**
 * 字符串转码
 */
func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

// 生成指定区间随机数
func Rand(from, to int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(to) + from
}

// Find获取一个切片并在其中查找元素。如果找到它，它将返回它的密钥，否则它将返回-1和一个错误的bool。
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// 组装钉钉消息
func AssemblyDing(title, text string) []byte {
	ding := make(map[string]interface{})
	ding["msgtype"] = "markdown"
	mark_down := make(map[string]string)
	mark_down["title"] = title
	mark_down["text"] = text
	ding["markdown"] = mark_down

	at := make(map[string]interface{})
	at["atMobiles"] = []string{}
	at["isAtAll"] = false
	ding["at"] = at

	ding_json, _ := json.Marshal(ding)
	return ding_json
}