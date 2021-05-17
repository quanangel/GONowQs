package routers

import (
	"nowqs/frame/http/blog/v1/controller"
	"nowqs/frame/http/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouters is admin routers function
func NewRouters(r *gin.Engine) *gin.Engine {

	v1Login := controller.NewLogin()
	blog := r.Group("/blog")
	{
		blog.Use(middleware.Logger("blog"))
		blog.Use(middleware.Cors())
		blog.GET("/v1/login", v1Login.Get)
		blog.PUT("/v1/login", v1Login.Put)
	}

	return r
}
