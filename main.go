package main

import (
	v1 "cfv-api/api/v1"
	"cfv-api/middleware"
	"fmt"

	"cfv-api/config"

	"github.com/gin-gonic/gin"
)

func main() {

	db, err := config.LoadDB()

	if err != nil {
		fmt.Println(err)
		return
	}

	app := gin.Default()
	api := app.Group("/api")
	api_v1 := api.Group("/v1")

	api_v1.Group("/v1").Use(middleware.SetDatabase(db))
	{
		api_v1.GET("/cards", middleware.SetDatabase(db), v1.GetCards)
	}

	app.Run()
}
