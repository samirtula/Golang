package main

import (
	"fmt"
	"reflect"
)

type myStructure struct {
	Name    string
	Surname string
	Height  int32
}

// конструктор
func NewStruct(n string, s string, h int32) *myStructure {
	if h > 300 {
		h = 0
	}

	return &myStructure{n, s, h}
}

func RetStructn(n string, s string, h int32) myStructure {
	if h > 300 {
		h = 0
	}

	return myStructure{n, s, h}
}

func main() {
	mike := NewStruct("Mike", "Smith", 88)
	fmt.Printf("Тип: %s\n", reflect.TypeOf(mike))

	mih := RetStructn("Mihalis", "Pipkin", 176)
	fmt.Printf("Тип: %s\n", reflect.TypeOf(mih))

	// Тип: *main.myStructure
	// Тип: main.myStructure

    fmt.Println((*mike).Name)
    fmt.Println(mih.Name)
    fmt.Println(mike)
    fmt.Println(mih)

	ps := new(myStructure)
	fmt.Println(ps)

	sP := new([]myStructure)
	fmt.Println(sP)
	*sP = make([]myStructure, 10)
	fmt.Println(sP)
}
