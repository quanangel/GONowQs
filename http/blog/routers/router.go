package routers

import (
	"nowqs/frame/http/blog/v1/controller"
	"nowqs/frame/http/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouters is blog routers function
func NewRouters(r *gin.Engine) {

	v1Login := controller.NewLogin()
	v1BlogClassify := controller.NewBlogClassify()
	v1Blog := controller.NewBlog()

	blog := r.Group("/blog")
	blog.Use(middleware.Logger("blog"))
	{
		// blog.Use(middleware.Logger("blog"))
		// blog.Use(middleware.Cors())
		blog.GET("/v1/login", v1Login.Get)
		blog.PUT("/v1/login", v1Login.Put)

		blog.GET("/v1/blog_classify", v1BlogClassify.Get)
		blog.POST("/v1/blog_classify", v1BlogClassify.Post)
		blog.PUT("/v1/blog_classify", v1BlogClassify.Put)
		blog.DELETE("/v1/blog_classify", v1BlogClassify.Delete)

		blog.GET("/v1/blog", v1Blog.Get)
		blog.POST("/v1/blog", v1Blog.Post)
		blog.PUT("/v1/blog", v1Blog.Put)
		blog.DELETE("/v1/blog", v1Blog.Delete)
	}
}
