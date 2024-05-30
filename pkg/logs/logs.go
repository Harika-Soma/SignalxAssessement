package logs

import (
	"log"
	"os"
	"time"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func CreateDir(folderName string) {
	_, err := os.Stat(folderName)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(folderName, 0755)
		if errDir != nil {
			log.Fatal("entered into dir", err)
		}
	}
}

func init() {
	CreateDir("supplychainlogs")
	filename := time.Now().Format("2006-01-02")
	InfoFile, err := os.OpenFile("supplychainlogs/Infolog-"+filename+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	WarningFile, err := os.OpenFile("supplychainlogs/Warninglog-"+filename+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	ErrorFile, err := os.OpenFile("supplychainlogs/Errorlog-"+filename+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(InfoFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(WarningFile, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(ErrorFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// func main() {
//     InfoLogger.Println("Starting the appliction.....")
//     InfoLogger.Println("Something noteworthy happened")
//     WarningLogger.Println("There is something you should know about")
//     ErrorLogger.Println("Something went wrong")
// }
