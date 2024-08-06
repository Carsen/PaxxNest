package ErrLog

import (
	"log"
	"os"
)

func LogErr(a error) {
	LogFile, err := os.OpenFile("errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer LogFile.Close()
	L := log.New(LogFile, "Error: ", log.Ldate|log.Ltime|log.Lshortfile)
	Logger(a, L)

}

func Logger(err error, logger *log.Logger) {
	if err != nil {
		logger.Print(err)
	}
}
