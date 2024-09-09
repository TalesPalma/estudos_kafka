package logservice

import (
	"fmt"
	"log"
	"os"
	"time"
)

func Logservice(msg string) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)

	if err != nil {
		log.Fatal("Error opening file", err)
	}
	defer file.Close()

	now := time.Now().Format("2006-01-02 15:04:05 PM")
	msgMounted := fmt.Sprintf("%s: %s\n", now, msg)
	file.WriteString(msgMounted)
}
