package databaseusers

import (
	"fmt"
	"log"
	"os"

	"github.com/TalesPalma/models"
)

func SaveUsersLogs(person models.Person) {
	file, err := os.OpenFile("./userLogs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)

	if err != nil {
		log.Fatal("error opening file:", err)
	}
	defer file.Close()

	formatingInfos := fmt.Sprintf("Persona : Name: %s , Age: %d", person.Name, person.Age)
	file.WriteString(formatingInfos + "\n")
}
