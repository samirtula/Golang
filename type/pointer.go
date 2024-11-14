package main

import (
	"fmt"
)

func getPointer(n *int) {
	*n = *n * *n
}

func returnPointer(n int) *int {
	v := n * n
	return &v
}

func main() {
	// & 
	// *
	x := 5
	getPointer(&x)
	fmt.Println(x)
	//25

	y := 7
	yPointer := returnPointer(y)
	getPointer(yPointer)
	fmt.Println("y", y)
	fmt.Println("yP", *yPointer)

    i := -10
	j := 25

	pI := &i
	pJ := &j

	fmt.Println("pI memory:", pI)
    fmt.Println("pI memory:", pJ)
    fmt.Println("pI value:", *pI)
    fmt.Println("pI value:", *pJ)
	//но. pI — это адрес памяти

	*pI = 123456
    *pI++
    fmt.Println("i:", i)
}