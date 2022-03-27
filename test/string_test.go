package test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestString1(t *testing.T) {
	s1 := "abc"
	ch := s1[0]

	fmt.Printf("type of s1: %T\n", s1)
	fmt.Printf("type of ch: %T\n", ch)

	fmt.Printf("ch: %d\n", ch)
	// fmt.Printf("ch: %s\n", ch)
	fmt.Printf("ch: %q\n", ch)

	//fmt.Printf("s1: %d\n", s1)
	fmt.Printf("s1: %s\n", s1)
	fmt.Printf("s1: %q\n", s1)

	fmt.Println(s1[0], s1[2])

}

func TestForLoop1(t *testing.T) {
	charCount := [2]int{0}
	for index, item := range charCount {
		fmt.Printf("type of index: %T\n", index)
		fmt.Printf("index: %d\n", index)
		fmt.Printf("type of item: %T\n", item)
		fmt.Printf("item: %d\n", item)
	}
}

func TestString2(t *testing.T) {
	x, _ := strconv.Atoi("123")             // x is an int
	y, _ := strconv.ParseInt("123", 10, 64) // base 10, up to 64 bits
	fmt.Printf("type x: %T, type y: %T\n", x, y)
	fmt.Printf("x=%d, y=%d\n", x, y)
}

func TestString3(t *testing.T) {
	// str1 := " aa bb ，,cc dd "

	str1 := "aa"
	str2 := strings.ReplaceAll(str1, "，", " ")
	str3 := strings.ReplaceAll(str2, ",", " ")

	strSlice := strings.Split(str3, " ")
	var finalStrSlice []string
	for _, name := range strSlice {
		if name != "" {
			finalStrSlice = append(finalStrSlice, name)
		}
	}
	finalStr := strings.Join(finalStrSlice, ",")
	t.Logf("str_slice: %+v", finalStr)
}
