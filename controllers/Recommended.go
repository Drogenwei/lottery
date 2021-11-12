package controllers

import (
	"lottery/common/helpers"
	"sort"
	"strconv"
	"strings"
)

/**
 * 双色球推荐
 */
func DoubleColor() string {
	len := 0
	str := "#### 双色球推荐 \n\n"
	for len < 3 {
		tui := tui()

		zhu := len + 1
		str = str + "**第" + strconv.Itoa(zhu)  +"注：**" + tui
		len++
	}

	return str
}

func tui() string {
	red_length := 6

	red := []string{}
	blue := helpers.Rand(1, 16)

	for red_length > 0 {
		rand := strconv.Itoa(helpers.Rand(1, 33))
		_, found := helpers.Find(red, rand)
		if found {
			continue;
		}
		red = append(red, rand)
		red_length--;
	}

	sort.Strings(red)
	red_str := strings.Join(red, " ")
	str := "红球：<font color=#FF0000>" + red_str + "</font> 篮球：<font color=#0000FF>" + strconv.Itoa(blue) + "</font> \n\n"
	return str
}