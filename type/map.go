package main

import "fmt"

func main() {
	//ловим панику
	defer func ()  {
		c := recover(); 

		if c != nil {
			fmt.Println("Panic inside", c)
		}
	}()

	iMap := make(map[string]int)
	// aMap := map[string]int{"k1": 12, "k2": 14}
	iMap["k1"] = 12
    iMap["k2"] = 13

	delete(iMap, "k1")

	_, ok := iMap["k3"]

	if ok {
		fmt.Println("Exists")
	} else {
		fmt.Println("Does Not exist")
	}

	for k, v := range iMap {
		fmt.Println(k, v)
	}

	aMap := map[string]int{}
	// var aMap map[string]int
	aMap = nil
	fmt.Println(aMap)
	//panic
	aMap["test"] = 1
}
