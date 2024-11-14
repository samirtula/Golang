package main

import (
	"fmt"
	"sort"
)

type aStructure struct {
	person string
	height int
	weight int
}

func main() {
	mySlice := make([]aStructure, 0)
	fmt.Println(mySlice)

	mySlice = append(mySlice, aStructure{"Ivan", 180, 99})
	mySlice = append(mySlice, aStructure{"Pupka", 160, 100})
	mySlice = append(mySlice, aStructure{"Khrupka", 175, 77})
    mySlice = append(mySlice, aStructure{"Epifanios", 144, 50})
	mySlice = append(mySlice, aStructure{"Marietta", 155, 45})

    fmt.Println("0:", mySlice)

	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].height < mySlice[j].height
	})

    fmt.Println("<:", mySlice)
    sort.Slice(mySlice, func(i, j int) bool {
        return mySlice[i].height > mySlice[j].height
    })
    fmt.Println(">:", mySlice)
}