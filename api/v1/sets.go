package v1

import (
	"cfv-api/constants"
	"cfv-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetSets returns the names of all sets as JSON
// TODO: Make queries case insensitive somehow?
// TODO: Enable querying by URL search params OR JSON
// depending on request header
func GetSets(c *gin.Context) {
	db := c.MustGet(constants.DB).(*gorm.DB)

	setQuery := models.Set{
		Name: c.Query("name"),
	}
	sets := []models.Set{}
	db.Model(&models.Set{}).
		Where(&setQuery).
		Find(&sets)
	c.JSON(200, sets)
}
