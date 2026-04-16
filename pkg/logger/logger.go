package logger

import (
	"log"
	"os"
)

var Log *log.Logger

func Init() {
	Log = log.New(os.Stdout, "[GOCOMMERCE] ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Info(msg string) {
	Log.Println("INFO:", msg)
}

func Error(msg string) {
	Log.Println("ERROR:", msg)
}