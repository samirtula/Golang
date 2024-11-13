package main

import "fmt"

func main() {
	//cрезы передаются в функции по ссылке
	//не указывается длина
	aSlice := []int{1, 2, 3, 4, 5, 6}
	//index out of range [9] with length 6
	// aSlice[9] = 1
	fmt.Println(aSlice)

	//необходимости автоматически расширяется
	integer := make([]int, 20)
	integer[19] = 1;
	fmt.Println(integer)
	//обязательно присвоение
	integer = append(integer, 2)
	fmt.Println(integer)
	aSlice2 := integer[0:3]
	fmt.Println(aSlice2)
	aSlice = nil
	//[]
	fmt.Println(aSlice)

	s1 := make([]int, 5)
    reSlice := s1[1:3]
    fmt.Println(s1)
    fmt.Println(reSlice)
    reSlice[0] = -100
    reSlice[1] = 123456
	//[0 -100 123456 0 0]
	// [-100 123456]
	//оба указываеют на один объект
    fmt.Println(s1)
    fmt.Println(reSlice)
	// при выходе за пределы capacity - capacity удваивается
}
