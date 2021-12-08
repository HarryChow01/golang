package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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

type HoloConfig struct {
	HoloHost     string `json:"holo_host"`
	HoloPort     int16  `json:"holo_port"`
	HoloUser string `json:"holo_user"`
	HoloPassword string `json:"holo_password"`
	HoloDBName   string `json:"holo_dbname"`
}

func loadJsonFile(filePath string) []byte {
	fmt.Printf("filePath: %s\n", filePath)
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fmt.Printf("content: %s\n", string(content))
	return content
}


func main() {
	filePath := "conf/conf.json"
	fileContent := loadJsonFile(filePath)

	holoConfig := HoloConfig{}
	err := json.Unmarshal(fileContent, &holoConfig)
	if err != nil {
		panic(err)
	}
	fmt.Printf("holoConfig: %+v\n", holoConfig)
}
