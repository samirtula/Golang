package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// получаем доступ ко всем элементам массива с помощью указателей.
	array := [...]int{0, 1, -2, 3, 4}
    pointer := &array[0]
    fmt.Print(*pointer, " ")
    memoryAddress := uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])
   
	for i := 0; i < len(array)-1; i++ {
        pointer = (*int)(unsafe.Pointer(memoryAddress))
        fmt.Print(*pointer, " ")
        memoryAddress = uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])
	}

	fmt.Println()
    pointer = (*int)(unsafe.Pointer(memoryAddress))
    fmt.Print("One more: ", *pointer, " ")
    memoryAddress = uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])
    fmt.Println()
}


func unsafePointer() {
	// разные типы
	var value int64 = 5
	var p1 = &value
	var p2 = (*int32)(unsafe.Pointer(p1))
	
	fmt.Println("*p1: ", *p1)
    fmt.Println("*p2: ", *p2)
    *p1 = 5434123412312431212
    fmt.Println(value)
    fmt.Println("*p2: ", *p2)
    *p1 = 54341234
    fmt.Println(value)
    fmt.Println("*p2: ", *p2)
}