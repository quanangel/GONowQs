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
	nav := adminController.NewNav()

	admin := r.Group("/admin")
	{
		admin.StaticFile("/swagger", "./http/admin/swagger/swagger.json")
		admin.GET("/login/index", login.Get)
		admin.PUT("/login/index", login.Put)

		admin.GET("/nav/index", nav.Get)
		admin.POST("/nav/index", nav.Post)
		admin.PUT("/nav/index", nav.Put)
	}

	return r
}
