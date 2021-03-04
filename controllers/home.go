package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Home routes to the homepage
func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{})
}

// TOS routes to the terms of service page with the given filename
func TOS(tosHTML string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, fmt.Sprintf("%v", tosHTML), gin.H{})
	}
}
