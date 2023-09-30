package main

import (
	"fmt"
	"math"
	"net/http"
	"reflect"
)

func hi(name string) string {
	temp := "temp"
	// Una variable booleana no declarada es falso por defecto
	// int8, int16, int32, int64
	// uint8, ...
	// float32, float64
	if len(name) > 10 {
		fmt.Println("long name" + temp)
	}
	var numbers [3]int
	numbers[2] = 10
	return "Hi " + name
}

func main() {
	fmt.Println(math.MaxInt64)

	// TypeOf
	fmt.Println(reflect.TypeOf(1))
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write([]byte("Hi"))
		if err != nil {
			return
		}
	})

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		return
	} else {
		fmt.Println("Server running...")
	}
}
