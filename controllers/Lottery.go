package controllers

import (
	"fmt"
	"lottery/common/helpers"
	"regexp"
)

// 红球匹配规则
var red_match = "<td class='WhiteBack RedFont' nowrap>(.*?)</td>"

// 篮球匹配规则
var blue_match = "<td class='WhiteBack BlueFont'>(.*?)</td>"

// 双色球开奖网站
var double_url = "https://zst.cjcp.com.cn/cjwssq/view/hong_zonghe_content.html"

var title = "双色球开奖通知"

var ding_url = "https://oapi.dingtalk.com/robot/send?access_token=e0716e0566710767aeb1944a20e5cd1198493a340aac3d3bc4de8dddbd01188d&sign="

// 双色球开奖信息获取
func Run() {
	html_str := helpers.Get(double_url)

	reg1 := regexp.MustCompile(red_match)
	if reg1 == nil { //解释失败，返回nil
		panic("regexp err")
	}
	result1 := reg1.FindAllStringSubmatch(html_str, -1)

	red := getLastElement(result1)
	fmt.Println(red)

	reg2 := regexp.MustCompile(blue_match)
	if reg2 == nil { //解释失败，返回nil
		panic("regexp err")
	}
	result2 := reg2.FindAllStringSubmatch(html_str, -1)
	blue := getLastElement(result2)
	fmt.Println(blue)

	// 获取markdown 格式文本
	markdown := getMarkdown(red, blue, "2021129")
	// 组装钉钉消息
	ding_info := helpers.AssemblyDing(title, markdown)
	// 发送
	helpers.SendPost(ding_url, ding_info)
}

// 获取markdown格式的发送文本
func getMarkdown(red, blue, nper string) string {
	str := "双色球第：**" + nper + "**期开奖：\n"
	str = str + "> **红球开奖结果：**<font color=#FF0000>" + red + "</font> \n\n"
	str = str + "> **蓝球开奖结果：**<font color=#0000FF>" + blue + "</font> \n\n"
	return str
}

// 获取切片的最后一个元素
func getLastElement(result [][]string) string {
	last_td := result[len(result) - 1]
	// 红球结果
	str := last_td[len(last_td) - 1]

	return str
}

