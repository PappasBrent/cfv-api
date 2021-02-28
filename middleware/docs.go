package middleware

import (
	"cfv-api/constants"
	"fmt"

	"github.com/gin-gonic/gin"
	redocMiddleware "github.com/go-openapi/runtime/middleware"
)

// GetRedocMiddleware returns a redoc middleware as a gin handlerFunc
func GetRedocMiddleware(apiVersion string) gin.HandlerFunc {
	redocOptions := redocMiddleware.RedocOpts{
		SpecURL:  constants.SwaggerPath,
		BasePath: fmt.Sprintf("/api/%v/", apiVersion),
	}
	sh := redocMiddleware.Redoc(redocOptions, nil)
	return gin.WrapH(sh)
}
