package models

import (
	"encoding/json"
	"log"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code        string `gorm:"unique;not null" json:"code"`
	Name        string `json:"name"`
	Price       string `json:"price"`
	Description string `json:"description"`
}

func (p *Product) IsEmpty() bool {
	if p.Code == "" || p.Name == "" || p.Price == "" || p.Description == "" {
		return true
	}
	return false
}

func (p *Product) MarshalJson() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		log.Fatalf("Erro ao empacotar os dados: %v", err)
	}
	return marshal
}
