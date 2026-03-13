package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string `json:"email" gorm:"unique;type:varchar(100)"`
	Password  string `json:"password" gorm:"type:varchar(100)"`
	FirstName string `json:"firstname" gorm:"type:varchar(100)"`
	LastName  string `json:"lastname" gorm:"type:varchar(100)"`
}
