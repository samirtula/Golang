package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	arguments := os.Args

	if len(arguments) < 2 {
	    fmt.Println("usage: switch number")
		os.Exit(1)
	}

	number, err := strconv.Atoi(arguments[1])

	if err != nil {
		fmt.Println("This value is not an integer:", number)
	} else {
		switch {
		case number < 0:
			fmt.Println("Less than zero!")
		case number > 0:
			fmt.Println("Bigger than zero!")
		default:
			fmt.Println("Zero")
		}
	}

	asString := arguments[1]
	var negative = regexp.MustCompile(`-`)
	var floatingPoint = regexp.MustCompile(`\d?\.\d`)
	var email = regexp.MustCompile(`^[^@]+@[^@.]+\.[^@.]+`)

	switch {
	case negative.MatchString(asString):
		fmt.Println("Negative number")
	case floatingPoint.MatchString(asString):
		fmt.Println("Floating point!")
	case email.MatchString(asString):
		fmt.Println("It is an email!")
	default:
	fmt.Println("Something else!")
	}
}