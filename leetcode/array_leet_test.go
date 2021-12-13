package leetcode

import (
	"fmt"
	"testing"
)

func GetFirstIndexArray1(data []int) (int, bool) {
	if len(data) == 0 {
		return -1, true
	}
	countMap := map[int]int{}
	for _, item := range data {
		countMap[item] += 1
	}
	for index, item := range data {
		if countMap[item] == 1 {
			return index, true
		}
	}
	return -1, true
}

func TestGetFirstIndexArray(t *testing.T) {
	data := []int{1, 2, 3, 1, 2}
	index, ok := GetFirstIndexArray1(data)
	fmt.Printf("ok: %t, index: %d\n", ok, index)
	data = []int{1, 2, 3}
	index, ok = GetFirstIndexArray1(data)
	fmt.Printf("ok: %t, index: %d\n", ok, index)
	data = []int{1, 2, 1, 2}
	index, ok = GetFirstIndexArray1(data)
	fmt.Printf("ok: %t, index: %d\n", ok, index)
}
