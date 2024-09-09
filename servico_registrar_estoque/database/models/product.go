package models

import (
	"encoding/json"
	"log"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID    int
	Name  string
	Price float64
	Stock int
	qty   int
}

func (p *Product) Unmarshal(data []byte) {
	err := json.Unmarshal(data, p)

	if err != nil {
		log.Fatalf("Erro with unmarshal: %v", err)
	}
}
