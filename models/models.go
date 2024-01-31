package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID      uint   `json:"id" gorm:"primaryKey"`
	Name    string `json:"name" gorm:"text;not null"`
	Surname string `json:"surname" gorm:"text default: null"`
}
