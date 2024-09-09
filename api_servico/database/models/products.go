package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code        string `gorm:"unique;not null" json:"code"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}

func (p *Product) IsEmpty() bool {
	if p.Code == "" || p.Name == "" || p.Price == 0 || p.Description == "" {
		return true
	}
	return false
}
