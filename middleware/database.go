package middleware

import (
	"cfv-api/constants"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetDatabase(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(constants.DB, db)
		c.Next()
	}
}
