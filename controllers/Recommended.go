package controllers

import (
	"lottery/common/helpers"
	"sort"
)

/**
 * 双色球推荐
 */
func DoubleColor() []int {
	red_length := 6

	red := []int{}
	blue := helpers.Rand(1, 16)

	for red_length > 0 {
		rand := helpers.Rand(1, 33)
		_, found := helpers.Find(red, rand)
		if found {
			continue;
		}
		red = append(red, rand)
		red_length--;
	}

	sort.Ints(red)
	res := append(red, blue)

	return res
}