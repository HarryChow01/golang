package test

import (
	"fmt"
	"testing"
)

func TestSlice1(t *testing.T) {
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
	var data2 []int;
	fmt.Println(data2)
	copy(data2, data1)
	fmt.Println(data2)

	data3 := make([]int, len(data1))
	fmt.Println(data3)
	copy(data3, data1)
	fmt.Println(data3)

}