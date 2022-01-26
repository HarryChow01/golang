package test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestRead1(t *testing.T) {
	file, err := os.Open("file1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, _ := ioutil.ReadAll(file)
	fmt.Println(string(content))
}

func TestRead2(t *testing.T) {
	filepath := "file1.txt"
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}


