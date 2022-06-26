package test

import (
	"fmt"
	"testing"
)

func TestSlice1(t *testing.T) {
	var s1 []int
	// s1 := []int(nil)  // 作用同上
	// s1 = nil		// 可以，但是无意义
	fmt.Printf("s1 == nil: %t\n", s1 == nil)
	fmt.Printf("len(s1): %d\n", len(s1))

	s2 := []int{}
	fmt.Printf("s2 == nil: %t\n", s2 == nil)
	fmt.Printf("len(s2): %d\n", len(s2))

}

func TestSlice5(t *testing.T) {
	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April",
		5: "May", 6: "June", 7: "July", 8: "August",
		9: "September", 10: "October", 11: "November", 12: "December"}
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)
	fmt.Println(summer)
	fmt.Println(cap(months))
	fmt.Println(cap(Q2))
}

func TestSliceCopy1(t *testing.T) {
	data1 := []int{1, 2, 3, 4, 5, 6, 7}
	var data2 []int
	fmt.Println(data2)
	copy(data2, data1)
	fmt.Println(data2)

	data3 := make([]int, len(data1))
	fmt.Println(data3)
	copy(data3, data1)
	fmt.Println(data3)

}
