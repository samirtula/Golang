package main

import "fmt"

func main() {
	anArray := [4]int{1, 2, -4, 4}
	fmt.Println(len(anArray))

	twoD := [4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	threeD := [2][2][2]int{{{1, 2}, {1, 2}}, {{1, 2}, {1, 2}}}

	for i:= 0; i <len(threeD); i++ {
		for j := 0; j < len(threeD[i]); j++ {
			for k:=0; k <len(threeD[i][j]); k++ {
				fmt.Println(threeD[i][j][k])
			}
		}
	}
	fmt.Println("----")
	
	for _, v := range twoD {
		for _, m := range v {
			fmt.Println(m)
		}
	}
}
