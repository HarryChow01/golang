
package main

import (
    "fmt"
    "unsafe"
)

func main() {
    // test_int()
    test_float()
}

func test_int() {
    var age1 int = 18
    var age2 int32
    var age3 = 0B1010
    fmt.Printf("age1: %d, age2: %d, age3: %d\n", age1, age2, age3)
    
    fmt.Println(unsafe.Sizeof(age3))
    
    var i1 int = 1
    var i2 int8 = 2
    var i3 int16 = 3
    var i4 int32 = 4
    var i5 int64 = 5
    fmt.Println(unsafe.Sizeof(i1))
    fmt.Println(unsafe.Sizeof(i2))
    fmt.Println(unsafe.Sizeof(i3))
    fmt.Println(unsafe.Sizeof(i4))
    fmt.Println(unsafe.Sizeof(i5))
}

func test_float() {
    var f1 float32 = 1.1
    var f2 float64 = 1.2
    var f3 = 1.3
    fmt.Println(unsafe.Sizeof(f1))
    fmt.Println(unsafe.Sizeof(f2))
    fmt.Println(unsafe.Sizeof(f3))
}


