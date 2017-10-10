package first

import (
	"fmt"
	"strconv"
)

const prefix= "first:"

func first(vars string) string {
	fmt.Println("first func called")
	name, age := "Poul", 40
	return vars + name + strconv.Itoa(age)
}
