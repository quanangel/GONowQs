package routers

import (
	adminController "nowqs/frame/http/admin/controller"
	adminMiddleware "nowqs/frame/http/admin/middleware"
	"nowqs/frame/http/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouters is admin routers function
func NewRouters(r *gin.Engine) *gin.Engine {
	r.Use(adminMiddleware.Logger())
	r.Use(middleware.Cors())
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	login := adminController.NewLogin()

	admin := r.Group("/admin")
	{
		admin.GET("/login/index", login.Get)
		admin.PUT("/login/index", login.Put)

	}

	return r
}
