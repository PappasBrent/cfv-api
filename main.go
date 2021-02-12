package main

import (
	"cfv-api/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {

	db, err := LoadDB()

	if err != nil {
		fmt.Println(err)
		return
	}

	app := gin.Default()

	app.GET("/", func(c *gin.Context) {
		if _, err := strconv.ParseInt(c.Query("id"), 10, 64); err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "invalid query parameter"})
		} else {
			cards := []models.Card{}
			db.Preload("Sets").Preload("TournamentStatuses").Where("id < 43 AND id > 41").Find(&cards)
			c.JSON(200, cards)
		}
	})

	app.Run()
}
