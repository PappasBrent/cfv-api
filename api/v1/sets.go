package v1

import (
	"cfv-api/constants"
	"cfv-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetSets returns the names of all sets as JSON
// TODO: Enable querying by URL search params OR JSON
// depending on request header
func GetSets(c *gin.Context) {
	db := c.MustGet(constants.DB).(*gorm.DB)

	sets := []models.Set{}
	db.Model(&models.Set{}).
		Find(&sets)
	c.JSON(200, sets)
}
