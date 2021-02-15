package config

import (
	v1 "cfv-api/api/v1"
	"cfv-api/controllers"
	"cfv-api/middleware"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
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

func SetupApp() (*gin.Engine, error) {
	db, err := LoadDB()

	if err != nil {
		fmt.Println(err)
		return nil, nil
	}

	app := gin.Default()
	app.LoadHTMLFiles("views/home.html")
	app.GET("/", controllers.Home)

	api := app.Group("/api")
	api_v1 := api.Group("/v1")
	api_v1.Group("/v1").Use(middleware.SetDatabase(db))
	{
		api_v1.GET("/cards", middleware.SetDatabase(db), v1.GetCards)
		api_v1.GET("/set", middleware.SetDatabase(db), v1.GetCardsInSet)
		api_v1.GET("/sets", middleware.SetDatabase(db), v1.GetSets)
	}

	return app, nil
}
