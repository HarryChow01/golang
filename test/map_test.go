package test

import (
	"fmt"
	"testing"
)

var mGlobal1 map[string]int
var mGlobal2 = map[string]int{
	"aa": 11,
	"bb": 22,
}

func TestMapGlobal1(t *testing.T) {
	fmt.Printf("mGlobal1 == nil: %t\n", mGlobal1 == nil)
	fmt.Printf("len(mGlobal1): %d\n", len(mGlobal1))

	fmt.Printf("mGlobal2 == nil: %t\n", mGlobal2 == nil)
	fmt.Printf("len(mGlobal2): %d\n", len(mGlobal2))
}

func TestMapLocal1(t *testing.T) {
	var mLocal map[string]int
	fmt.Printf("mLocal == nil: %t\n", mLocal == nil)
	fmt.Printf("len(mLocal): %d\n", len(mLocal))

	// ok
	m1 := map[string]int{}
	m2 := map[string]int{
		"aa": 11,
		"bb": 22,
	}
	fmt.Printf("m1 == nil: %t\n", m1 == nil)
	fmt.Printf("m2 == nil: %t\n", m2 == nil)
	fmt.Printf("len(m1): %d\n", len(m1))
	fmt.Printf("len(m2): %d\n", len(m2))

	m3 := make(map[string]int)
	fmt.Printf("m3 == nil: %t\n", m3 == nil)
	fmt.Printf("len(m3): %d\n", len(m3))
}

func TestMap1(t *testing.T) {
	map1 := make(map[string]int)
	map1["aa"] = 11
	fmt.Println(map1)

	map2 := map[string]int{
		"aa": 11,
		"bb": 22,
	}
	fmt.Println(map2)

	map2["aa"] += 1
	map2["bb"]++
	fmt.Println(map2)

	for key, value := range map2 {
		fmt.Printf("%s\t%d\n", key, value)
	}

	for key := range map2 {
		fmt.Printf("%s\t%d\n", key, map2[key])
	}

	// 判断某个Key是否存在
	if _, ok := map2["Tony"]; ok {
		fmt.Println("Tony is exists")
	}

	if value, ok := map2["aa"]; ok {
		fmt.Printf("value: %d\n", value)
	} else {
		fmt.Printf("value: %d\n", value)
	}

}
