package main

import (
	"fmt"
	"errors"
	"os"
)

func main() {
	err := returnErrors(1, 2)
	if err == nil {
		fmt.Println("returnError() ended normally!")
	} else {
		fmt.Println(err)
	}

	err = returnErrors(2, 2)
	if err == nil {
		fmt.Println("returnError() ended normally!")
	} else {
		fmt.Println(err)
	}
	// возвращает string
	if err.Error() == "Error in returnError() function!" {
		panic(err)
		os.Exit(10)
	}
}

func returnErrors(a, b int) error {
	if a == b {
		err := errors.New("Error in returnError() function!")
		return err
	}

	return nil
}