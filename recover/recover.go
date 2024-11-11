package main

import (
	"fmt"
)

func a() {
	fmt.Println("Inside func a")

	defer func() {
		c := recover(); 

		if c != nil {
			fmt.Println("Recover inside a()!")
		}
	}()

	fmt.Println("About to call func b")
	b()
	fmt.Println("func b bexited!")
	fmt.Println("Exiting func a")
}

func b() {
	fmt.Println("Inside func b")
	panic("Panic in func b!")
	fmt.Println("Exiting func b")
}

func main() {
	a()
	fmt.Println("func main ended")

// 	не выполнились 
// fmt.Println("func b bexited!")
// fmt.Println("Exiting func a")
// defer + recover оработал и вышел из программы
}