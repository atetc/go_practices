package main

import (
	"fmt"
	"strconv"
)

var name string
var age int

func first(vars string) string {
	fmt.Println("first func called")
	name, age := "Poul", 40
	return vars + name + strconv.Itoa(age)
}

func main() {
	fmt.Println(first("first:"))
}