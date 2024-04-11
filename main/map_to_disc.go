package main

import (
	"fmt"
	"math/rand"
	"os"
	"reflect"
)

func main() {
	filepath := "ressources/plan-brussels.jpg"
	data, err := os.ReadFile(filepath)
	fmt.Println(len(data))
	fmt.Println(err)
	fmt.Println(reflect.TypeOf(data))
	for i := 5000; i < len(data); i += 50000 {
		data[i] = byte(rand.Intn(200))
	}
	os.WriteFile("plan2.jpg", data, os.ModePerm)
}
