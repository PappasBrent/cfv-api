package config

import (
	v1 "cfv-api/api/v1"
	"cfv-api/controllers"
	"cfv-api/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

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
	app.LoadHTMLGlob("views/*.html")

	app.GET("/", controllers.Home)

	api := app.Group("/api")
	apiV1 := api.Group("/v1")

	apiV1.Use(middleware.SetDatabase(db))
	{
		apiV1.GET("/card", v1.GetCard)
		apiV1.GET("/cards", v1.GetCards)
		apiV1.GET("/set", v1.GetCardsInSet)
		apiV1.GET("/sets", v1.GetSets)
	}
	apiV1.GET("/docs", middleware.GetRedocMiddleware("v1"))
	apiV1.GET("/tos", controllers.TOS("tosv1.html"))

	return app, nil
}
