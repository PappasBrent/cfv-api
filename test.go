package main

import (
	"os"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestLoadDSN(t *testing.T) {

	if _, err := LoadDSN(); err != nil {
		t.Error(err)
	}

	envKeys := []string{"DB_HOST", "DB_USER", "DB_PASSWORD", "DB_DB", "DB_PORT"}

	for _, key := range envKeys {
		if _, exists := os.LookupEnv(key); exists == false {
			t.Errorf("Error looking up environment variable %s", key)
		}
	}
}

func TestConnection(t *testing.T) {

	dsn, err := LoadDSN()

	if err != nil {
		t.Error(err)
	}

	dialector := postgres.New(postgres.Config{DSN: dsn, PreferSimpleProtocol: true})

	if _, err := gorm.Open(dialector, &gorm.Config{}); err != nil {
		t.Error(err)
	}
}
