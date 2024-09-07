package logservice

import (
	"log"
	"os"
)

func Logservice(msg string) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)

	if err != nil {
		log.Fatal("Error opening file", err)
	}
	defer file.Close()

	file.WriteString(msg + "\n")
}
