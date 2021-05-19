package routers

import (
	"nowqs/frame/http/blog/v1/controller"
	"nowqs/frame/http/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouters is admin routers function
func NewRouters(r *gin.Engine) *gin.Engine {

	v1Login := controller.NewLogin()
	v1BlogClassify := controller.NewBlogClassify()
	blog := r.Group("/blog")
	{
		blog.Use(middleware.Logger("blog"))
		blog.Use(middleware.Cors())
		blog.GET("/v1/login", v1Login.Get)
		blog.PUT("/v1/login", v1Login.Put)

		blog.GET("/v1/blog_classify", v1BlogClassify.Get)

	}

	return r
}
