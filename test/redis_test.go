package test

import (
    "fmt"
    "path"
    "runtime"
    "testing"
)

func TestCurrentAbPathByCaller(t *testing.T) {
    var abPath string
    _, filename, _, ok := runtime.Caller(0)
    fmt.Printf("cur file: %s\n", filename)
    if ok {
        abPath = path.Dir(filename)
    }
    fmt.Printf("absolute path: %s\n", abPath)
}
