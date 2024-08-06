package ErrLog

import (
	"log"
	"os"
)

func LogErr(a error) {
	LogFile, err := os.OpenFile("errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	defer LogFile.Close()
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	log.SetOutput(LogFile)
	log.Print(a)
}
