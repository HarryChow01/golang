package test

import (
	"fmt"
	"testing"
)

type Point struct{ X, Y int }

func TestStruct1(t *testing.T) {
	pp := &Point{1, 2}	//
	fmt.Println(&pp)

	qq := new(Point)	//
	fmt.Println(&qq)
	*qq = Point{1, 2}
	fmt.Println(&qq)
}