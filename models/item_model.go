package models

import "gorm.io/gorm"

type ItemModel struct {
	gorm.Model
	name string
	amount uint
}
