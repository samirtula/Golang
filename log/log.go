package main

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
	"path/filepath"
)

func main() {
	logInfo()
	logMail()
	customLog()
}

func logMail() {
	sysLog, err := syslog.New(syslog.LOG_MAIL, "")

	if err != nil {
		// log.Panic( выводит дополнительную низкоуровневую информацию
		// log.Fatal(), функция log.Panic() добавит запись в соответствующий файл журнала и немедленно прекратит работу Go-программы
		log.Fatal(err)
	} else {
		log.SetOutput(sysLog)
	}

	log.Println("LOG_MAIL: Logging in Go!")
	fmt.Println("Will you see this?")
}

func logInfo() {
	programName := filepath.Base(os.Args[0])
	sysLog, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, programName)

	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sysLog)
	}

	log.Println("LOG_INFO + LOG_LOCAL7: Logging in Go!")
}

func customLog() {
	var LOGFILE = "/tmp/mGo.log"
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	iLog := log.New(f, "customLogLineNumber ", log.LstdFlags)

	iLog.SetFlags(log.LstdFlags)
	iLog.Println("Hello there!")
	iLog.Println("Another log entry!")
}
