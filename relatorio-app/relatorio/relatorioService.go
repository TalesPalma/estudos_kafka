package relatorio

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func GenerateRelatorio(msg []byte) error {

	dataFormat := time.Now().Format("2006-01-02")
	fileName := fmt.Sprintf("relatorio_%s.txt", dataFormat)
	pathFile := filepath.Join("relatorios", fileName)
	file, err := os.Create(pathFile)

	if err != nil {
		return err
	}
	defer file.Close()

	file.Write(msg)

	return nil
}
