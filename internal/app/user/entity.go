package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string
	Username string
	Password string
}
