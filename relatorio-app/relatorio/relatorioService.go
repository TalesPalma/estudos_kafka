package relatorio

import (
	"fmt"
	"os"
	"time"
)

func GenerateRelatorio(msg []byte) error {

	dataFormat := time.Now().Format("2006-01-02")
	file, err := os.Create(fmt.Sprintf("relatorio_%s.txt", dataFormat))

	if err != nil {
		return err
	}
	defer file.Close()

	file.Write(msg)

	return nil
}
