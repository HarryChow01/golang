package test

import (
	"fmt"
	"testing"
)

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

}