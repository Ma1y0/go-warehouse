package models

import "gorm.io/gorm"

type ItemModel struct {
	gorm.Model
	Name string
	Amount uint 
}
