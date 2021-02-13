package v1

import (
	"cfv-api/constants"
	"cfv-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetCards(c *gin.Context) {
	db := c.MustGet(constants.DB).(*gorm.DB)
	name := c.Query("name")
	cards := []models.Card{}
	db.Preload("Sets").
		Preload("TournamentStatuses").
		Where(`name = ?`, name).
		Find(&cards)
	c.JSON(200, cards)
}
