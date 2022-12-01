package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"size:225;not null; unique" json:"username"`
	Password string `gorm:"size:225;not null;" json:"password"`
}