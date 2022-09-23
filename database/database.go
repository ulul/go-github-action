package database

import (
	"fmt"
	"hexagonal-architecture/internal/app/book"
	"hexagonal-architecture/internal/app/user"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"))

	var e error
	DB, e = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if e != nil {
		panic(e)
	}
	Migrate()
}

func Migrate() {
	DB.AutoMigrate(&user.User{})
	DB.AutoMigrate(&book.Book{})
}
