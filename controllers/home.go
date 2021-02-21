package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Home routes to the homepage
func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{})
}
