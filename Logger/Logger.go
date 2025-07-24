package logger

import (
	"log"
	"os"
)

var (
	LogFile *os.File
	Logger  *log.Logger
)

func InitLogger() {
	var err error
	LogFile, err = os.OpenFile("qhospital.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	Logger = log.New(LogFile, "", log.LstdFlags)
}

func CloseLogger() {
	LogFile.Close()
}

func Info(msg string) {
	Logger.Println("[INFO] " + msg)
}

func Error(msg string) {
	Logger.Println("[ERROR] " + msg)
}
