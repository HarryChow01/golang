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

// 给定两个由小到大有序int型数组，返回1个归并后由小到大有序的int型数组
func merge(a, b []int) []int {
	i, j := 0, 0
	var res []int
	for {
		if a[i] < b[j] {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			j++
		}
		if i == len(a) || j == len(b) {
			break
		}
	}
	if i == len(a) {
		res = append(res, b[j:]...)
	} else {
		res = append(res, a[i:]...)
	}
	return res
}
