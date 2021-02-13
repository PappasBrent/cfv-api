package v1

import (
	"cfv-api/constants"
	"cfv-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetCards(c *gin.Context) {
	db := c.MustGet(constants.DB).(*gorm.DB)
	if _, err := strconv.ParseInt(c.Query("id"), 10, 64); err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "invalid query parameter"})
	} else {
		cards := []models.Card{}
		db.Preload("Sets").Preload("TournamentStatuses").Where("id < 43 AND id > 41").Find(&cards)
		c.JSON(200, cards)
	}
}
