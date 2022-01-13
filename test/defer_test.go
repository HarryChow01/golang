package test

import (
	"fmt"
	"testing"
)

func testHelper3() bool {
	defer fmt.Println("in defer testHelper3")
	fmt.Println("test testHelper3")
	return true
}

func testHelper1() bool {
	defer fmt.Println("in defer testHelper1")
	fmt.Println("test testHelper1")
	return testHelper3()
}

func testHelper2() {
	testHelper1()
	fmt.Println("test testHelper2")
}

func TestDefer1(t *testing.T) {
	testHelper2()
}
