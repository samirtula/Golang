package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

func main() {
	fmt.Println("Epoch time:", time.Now().Unix())
	t := time.Now()
	fmt.Println(t, t.Format(time.RFC3339))
	fmt.Println(t.Weekday(), t.Day(), t.Month(), t.Year())

	time.Sleep(time.Second)
	t1 := time.Now()
	fmt.Println("Time difference:", t1.Sub(t))

	//создаем формат
	formatT := t.Format("01 Jan 2006")
	fmt.Println(formatT)

	loc, _ := time.LoadLocation("Asia/Baku")
	bakuTime := t.In(loc)
	fmt.Println("Baku:", bakuTime)

	var myDate string

	if len(os.Args) != 2 {
		fmt.Printf("usage: %s string\n", filepath.Base(os.Args[0]))
		return
	}

	myDate = os.Args[1]
	d, err := time.Parse("02 January 2006", myDate)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Time:", d.Day(), d.Month(), d.Year())
	} else {
		fmt.Println(err)
	}

	logs := []string{"127.0.0.1 - - [16/Nov/2017:10:49:46 +0200] 325504",
		"127.0.0.1 - - [16/Nov/2017:10:16:41 +0200]"}

	for _, logEntry := range logs {
		r := regexp.MustCompile(`.*\[(\d\d\/\w+/\d\d\d\d:\d\d:\d\d:\d\d.*)\].*`)
		
		if r.MatchString(logEntry) {
			match := r.FindStringSubmatch(logEntry)
			dt, err := time.Parse("02/Jan/2006:15:04:05 -0700", match[1])
			
			if err == nil {
				newFormat := dt.Format(time.RFC850)
				fmt.Println(newFormat)
			} else {
				fmt.Println("Not a valid date time format!")
			}
		} else {
			fmt.Println("Not a match!")
		}
	}
}
