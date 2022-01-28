package test

import (
	"fmt"
	"reflect"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

func TestBoolFmt1(t *testing.T) {
	v1 := true
	v2 := false
	fmt.Printf("%t\n", v1)
	fmt.Printf("%t\n", v2)
}

func TestIntFmt1(t *testing.T) {
	intValue1 := 15
	fmt.Printf("%d\n", intValue1)
	fmt.Printf("%b\n", intValue1)

	fmt.Printf("%x\n", intValue1)
	fmt.Printf("%X\n", intValue1)

	fmt.Printf("%o\n", intValue1)
	fmt.Printf("%O\n", intValue1)

}

func TestGeneralFmt1(t *testing.T) {
	person := Person{
		Name: "Harry",
		Age:  30,
	}

	fmt.Printf("%v\n", person)
	fmt.Printf("%+v\n", person)
	fmt.Printf("%#v\n", person)
	fmt.Printf("%%\n")

	fmt.Printf("%T\n", person)
	fmt.Printf("%s\n", reflect.TypeOf(person))
	fmt.Printf("%s\n", reflect.TypeOf(person).Name())
	fmt.Printf("%s\n", reflect.TypeOf(person).Kind())

	var v1 int64 = 5
	fmt.Printf("%v\n", v1)
	fmt.Printf("%+v\n", v1)
	fmt.Printf("%#v\n", v1)
	fmt.Printf("%T\n", v1)
	fmt.Printf("%s\n", reflect.TypeOf(v1))
	fmt.Printf("%s\n", reflect.TypeOf(v1).Name())
	fmt.Printf("%s\n", reflect.TypeOf(v1).Kind())

	var v2 float32 = 1.2
	fmt.Printf("%v\n", v2)
	fmt.Printf("%+v\n", v2)
	fmt.Printf("%#v\n", v2)
	fmt.Printf("%T\n", v2)
	fmt.Printf("%s\n", reflect.TypeOf(v2))
	fmt.Printf("%s\n", reflect.TypeOf(v2).Name())
	fmt.Printf("%s\n", reflect.TypeOf(v2).Kind())
}

func TestVar1(t *testing.T) {
	var aa int = 10
	fmt.Printf("aa: %+v\n", aa)
}
