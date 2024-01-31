package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name    string `json:"name" gorm:"text;not null;default:null`
	Surname string `json:"surname" gorm:"text;default:null`
}
