package util

import (
	"log"
	"os"
)

type Logger struct {
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
}

func NewLogger() (*Logger, error) {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	return &Logger{
		InfoLogger:    log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		WarningLogger: log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile),
		ErrorLogger:   log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}, nil
}
