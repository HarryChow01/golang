package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name    string
	Age     int
	Failed   bool
	Classes []string
	Price   float32
}

func jsonTest1() {
	student1 := &Student {
		"Harry",
		25,
		false,
		[]string{"Math", "English", "Chinese"},
		9.99,
	}

	fmt.Println("student1:", student1)
	jsonBytes, _ := json.Marshal(student1)
	// fmt.Println("jsonStr:", jsonStr)		// byte
	fmt.Println("jsonStr:", string(jsonBytes))	// string

	student2 := new(Student)
	fmt.Println("student2:", student2)
	json.Unmarshal(jsonBytes, &student2)
	// json.Unmarshal([]byte(string(jsonBytes)), &student2)
	fmt.Println("student2:", student2)
}


func main() {
	jsonTest1()
}
