package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 100; i++ {
	}

	i := 0
	anExpression := true
	for ok := true; ok; ok = anExpression {
		if i > 10 {
			anExpression = false
		}
		fmt.Println(i, " ")
		i++
	}
	anArray := [5]int{0, -1, -2}

	for i, value := range anArray {
		fmt.Println("index:", i, "value: ", value)
	}
}
