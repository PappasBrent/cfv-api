// Package v1 classification of Sets API
//
// Documentation for Sets API
//
// Schemes: http
// BasePath: /api/v1/
// Version: 1.0.0
//
// Produces:
// - application/json
// swagger:meta
package v1

import (
	"cfv-api/constants"
	"cfv-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// swagger:route GET /sets sets getSets
// Returns all sets
// responses:
//  200: setsResponse

// A list of all set names
// swagger:response setsResponse
type setsResponse struct {
	// A list of all set names
	// in: body
	Body []string
}

// GetSets returns the names of all sets as JSON
// swagger:meta
func GetSets(c *gin.Context) {
	db := c.MustGet(constants.DB).(*gorm.DB)

	sets := []models.Set{}
	db.Model(&models.Set{}).
		Find(&sets)
	c.JSON(200, sets)
}
