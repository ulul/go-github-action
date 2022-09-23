package main

import (
	"hexagonal-architecture/database"
	"hexagonal-architecture/internal/routes"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.InitDB()

	e := routes.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Logger.Fatal(e.Start(":8080"))
}
