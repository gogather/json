package json

import (
	"log"
	"testing"
)

func Test_Marshal(t *testing.T) {
	// string
	val1 := "string"
	result, err := Marshal(val1)
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println("result: " + string(result))
	}

	// array
	val2 := []int{1, 2, 3, 4, 5, 6}
	result, err = Marshal(val2)
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println("result: " + string(result))
	}

	// map
	val3 := map[string]interface{}{
		"hello": "world",
		"me":    "we",
	}
	result, err = Marshal(val3)
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println("result: " + string(result))
	}

	// struct
	type User struct {
		Id   int    `json:"ID"`
		Name string `json:"nAmE"`
		Age  uint   `json:"AGe"`
	}
	val4 := &User{
		Id:   1,
		Name: "Rex",
		Age:  28,
	}
	result, err = Marshal(val4)
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println("result: " + string(result))
	}

}
