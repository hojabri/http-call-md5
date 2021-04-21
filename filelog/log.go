package filelog

import (
	"log"
	"os"
)
var Logger *log.Logger
var file *os.File

// Custom log, for printing logs to file

func init() {

	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	Logger = log.New(file, "LOG: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Close() {
	file.Close()
}