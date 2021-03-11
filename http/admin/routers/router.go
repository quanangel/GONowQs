package routers

import (
	adminMiddleware "nowqs/frame/http/admin/middleware"
	"nowqs/frame/http/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouters is admin routers function
func NewRouters(r *gin.Engine) *gin.Engine {
	r.Use(adminMiddleware.Logger())
	r.Use(middleware.Cors())
	return r
}
