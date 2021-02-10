package main

import (
	"fmt"

	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadDSN() (string, error) {
	// TODO: Check environment (dev vs. prod) before loading dotenv.
	// Store that as a parameter?
	if err := godotenv.Load(); err != nil {
		return "", err
	}

	host, user, password, db, port := os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DB"), os.Getenv("DB_PORT")
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, db, port), nil
}

func main() {

	dsn, err := LoadDSN()

	if err != nil {
		fmt.Println(err)
		return
	}

	dialector := postgres.New(postgres.Config{DSN: dsn, PreferSimpleProtocol: true})
	_, err = gorm.Open(dialector, &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Success!")

}
