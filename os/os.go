package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	// stdOut()
	// stdIn()
	fmt.Println(minMax())
}

func stdOut() {
	myString := ""
	//аргумент командной строки
	args := os.Args

	if len(args) == 1 {
		myString = "Please give me one argument"
	} else {
		myString = args[1]
	}

	io.WriteString(os.Stdout, myString)
	io.WriteString(os.Stdout, "\n")
}

func stdIn() {
	var f *os.File
	//стандартный поток ввода	
	f = os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
	}
}

func minMax() (float64, float64) {
	if len(os.Args) < 2 {
		fmt.Println("Please give one or more floats.")
		os.Exit(1)
	}

	args := os.Args
	min, _ := strconv.ParseFloat(args[1], 64)
	max, _ := strconv.ParseFloat(args[1], 64)

	for i := 2; i < len(args); i++ {
		n, _ := strconv.ParseFloat(args[i], 64)

		if n < min {
			min = n
		}
		
		if n > max {
			max = n
		}
	}

	return min, max;
}
