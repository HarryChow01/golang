
package main

import (
    "fmt"
    "unsafe"
)

func main() {
    // test_int()
    // test_float()
    // test_bool()
    // test_addr()
    test_data()
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

func test_bool() {
    var xx = true
    var yy = false
    fmt.Println(unsafe.Sizeof(xx))
    fmt.Println(unsafe.Sizeof(yy))
}

func test_addr() {
    var a int = 10  
    fmt.Printf("address of a: %x\n", &a)
    
    var ptr *int
    fmt.Printf("ptr : %x\n", ptr)
    // var ip *int        /* 指向整型*/
    // var fp *float32    /* 指向浮点型 */
    // if(ptr != nil)     /* ptr is null */
    // if(ptr == nil)    /* ptr is not null */
}

func test_data() {
    var x uint8 = 1<<1 | 1<<5
    var y uint8 = 1<<1 | 1<<2

    fmt.Printf("%08b\n", x) // "00100010", the set {1, 5}
    fmt.Printf("%08b\n", y) // "00000110", the set {1, 2}

    fmt.Printf("%08b\n", x&y)  // "00000010", the intersection {1}
    fmt.Printf("%08b\n", x|y)  // "00100110", the union {1, 2, 5}
    fmt.Printf("%08b\n", x^y)  // "00100100", the symmetric difference {2, 5}
    fmt.Printf("%08b\n", x&^y) // "00100000", the difference {5}

    for i := uint(0); i < 8; i++ {
        if x&(1<<i) != 0 { // membership test
            fmt.Println(i) // "1", "5"
        }
    }

    fmt.Printf("%08b\n", x<<1) // "01000100", the set {2, 6}
    fmt.Printf("%08b\n", x>>1) // "00010001", the set {0, 4}
}
























