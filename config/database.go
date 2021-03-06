package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// LoadDSN loads the data source name for connecting to the database.
// It checks the GIN_MODE environment variable to determine whether
// to load the DSN from the environment or a .env file
func LoadDSN() (string, error) {
	if mode := os.Getenv(("GIN_MODE")); mode != "RELEASE" {
		if err := godotenv.Load(); err != nil {
			return "", err
		}
	}

	dbFilepath, mode := os.Getenv("DB_NAME"), os.Getenv("MODE")
	return fmt.Sprintf("file:%s?mode=%s", dbFilepath, mode), nil
}

// LoadDB loads the gorm DB
func LoadDB() (*gorm.DB, error) {
	dsn, err := LoadDSN()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	dialector := sqlite.Open(dsn)
	db, err := gorm.Open(dialector, &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return db, nil
}
