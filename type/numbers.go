package main

import (
	"fmt"
)

func main() {
	x := 12
	k := 5

	fmt.Println(x)
	fmt.Printf("Type of x: %T\n", x)

	// div := x / k
	// Некорректный результат div == 2
	// fmt.Println("div", div)

	var m, n float64
	m = 1.223
	fmt.Println("m, n:", m, n)

	y := 4 / 2.3
	//1.7391304347826086
	fmt.Println("y:", y)

	divFloat := float64(x) / float64(k)
	//2.4
	fmt.Println("divFloat", divFloat)
	fmt.Printf("Type of divFloat: %T\n", divFloat)
}

func complexNums() {
	c1 := 12 + 1i
	c2 := complex(5, 7)

	fmt.Printf("Type of c1: %T\n ", c1)
	fmt.Printf("Type of c2: %T\n ", c2)

	var c3 complex64 = complex64(c1 + c2)
	fmt.Println("c3:", c3)
	fmt.Printf("Type of c3: %T\n ", c3)
	cZero := c3 - c3
	fmt.Println("cZero:", cZero)
}
