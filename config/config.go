package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadDSN() (string, error) {
	// Load the Data source name for connecting to the database.
	// Checks the GIN_MODE environment variable to determine whether
	// to load the DSN from the environment or a .env file
	if mode := os.Getenv(("GIN_MODE")); mode != "RELEASE" {
		if err := godotenv.Load(); err != nil {
			return "", err
		}
	}

	host, user, password, db, port := os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DB"), os.Getenv("DB_PORT")
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, db, port), nil
}

func LoadDB() (*gorm.DB, error) {
	// Loads the gorm DB
	dsn, err := LoadDSN()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	dialector := postgres.New(postgres.Config{DSN: dsn, PreferSimpleProtocol: true})
	db, err := gorm.Open(dialector, &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return db, nil
}
