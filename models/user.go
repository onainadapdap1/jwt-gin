package models

import (
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:225;not null; unique" json:"username"`
	Password string `gorm:"size:225;not null;" json:"password"`
}

// set method SaveUser milik struct User
func (u *User) SaveUser() (*User, error) {
	if err := DB.Create(&u).Error; err != nil {
		return &User{}, err
	}
	return u, nil
}

// set method BeforeSave milik struct User, turn password into hash
func (u *User) BeforeSave() error {
	// turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare() {
	u.ID = 0
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}