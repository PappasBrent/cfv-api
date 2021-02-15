package v1

import (
	"cfv-api/constants"
	"cfv-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// TODO: Enable querying by URL search params OR JSON
// depending on request header
// TODO: Make queries case insensitive somehow?
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
