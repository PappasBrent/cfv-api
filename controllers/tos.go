package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TOS routes to the terms of service page with the given filename
func TOS(tosHTML string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, fmt.Sprintf("%v", tosHTML), gin.H{})
	}
}
