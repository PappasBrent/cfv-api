package config

import (
	v1 "cfv-api/api/v1"
	"cfv-api/controllers"
	"cfv-api/middleware"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
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

// SetupApp creates the gin router, attaches the gorm GB to it,
// and and sets up its routes
func SetupApp() (*gin.Engine, error) {
	db, err := LoadDB()

	if err != nil {
		fmt.Println(err)
		return nil, nil
	}

	app := gin.Default()
	app.Static("/assets", "./assets")
	app.LoadHTMLFiles("views/home.html")

	app.GET("/", controllers.Home)

	api := app.Group("/api")
	apiV1 := api.Group("/v1")

	apiV1.Group("/v1").Use(middleware.SetDatabase(db))
	{
		apiV1.GET("/cards", middleware.SetDatabase(db), v1.GetCards)
		apiV1.GET("/set", middleware.SetDatabase(db), v1.GetCardsInSet)
		apiV1.GET("/sets", middleware.SetDatabase(db), v1.GetSets)
	}

	return app, nil
}
