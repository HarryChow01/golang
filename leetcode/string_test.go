package leetcode

import (
	"fmt"
	"testing"
)

func TestString1(t *testing.T) {
	s1 := "abc"
	ch := s1[0]

	fmt.Printf("type of s1: %T\n", s1)
	fmt.Printf("type of ch: %T\n", ch)

	fmt.Printf("ch: %d\n", ch)
	fmt.Printf("ch: %q\n", ch)
}

func TestForLoop1(t *testing.T) {
	charCount := [26]int{0}
	for index, item := range charCount {
		fmt.Printf("type of index: %T\n", index)
		fmt.Printf("index: %d\n", index)
		fmt.Printf("type of item: %T\n", item)
		fmt.Printf("item: %d\n", item)
	}
}

func GetFirstIndex1(str string) (int, bool) {
	if str == "" {
		return -1, true
	}

	charCount := [26]int{0}
	for _, ch := range str {
		// fmt.Printf("type of ch: %T\n", ch)
		pos := ch - 'a'
		if pos < 0 || pos >= 26 {
			fmt.Printf("ch: %q error", ch)
			return -1, false
		}
		charCount[pos] += 1
	}

	for index, ch := range str {
		pos := ch - 'a'
		if charCount[pos] == 1 {
			return index, true
		}
	}
	return -1, true
}

func GetFirstIndex2(str string) (int, bool) {
	if str == "" {
		return -1, true
	}

	charCountMap := map[int32]int{}
	for _, ch := range str {
		if ch < 'a' || ch > 'z' {
			fmt.Printf("ch: %q error", ch)
			return -1, false
		}
		charCountMap[ch] += 1
	}

	for index, ch := range str {
		if charCountMap[ch] == 1 {
			return index, true
		}
	}
	return -1, true
}

func TestGetFirstIndex(t *testing.T) {
	s1 := "abcd"
	index, ok := GetFirstIndex2(s1)
	fmt.Printf("ok: %t, index: %d\n", ok, index)
	s2 := "abcab"
	index, ok = GetFirstIndex2(s2)
	fmt.Printf("ok: %t, index: %d\n", ok, index)
	s3 := "ababcc"
	index, ok = GetFirstIndex2(s3)
	fmt.Printf("ok: %t, index: %d\n", ok, index)
}
