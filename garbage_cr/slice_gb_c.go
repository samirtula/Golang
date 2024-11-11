package main

import (
	"runtime"
)

type data struct {
	i, j int
}

func main() {
	var N = 40000000
	var structure []data

	for i := 0; i < N; i++ {
		value := int(i)
		structure = append(structure, data{value, value})
	}

	runtime.GC()
	//Последний оператор (_ = structure[0]) используется для предотвращения пре- ждевременной очистки сборщиком мусора переменной structure
	_ = structure[0]
}
