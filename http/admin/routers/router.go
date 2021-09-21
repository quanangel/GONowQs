package routers

import (
	adminController "nowqs/frame/http/admin/controller"
	"nowqs/frame/http/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouters is admin routers function
func NewRouters(r *gin.Engine) {

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	login := adminController.NewLogin()
	nav := adminController.NewNav()
	admin := r.Group("/admin")
	{
		admin.Use(middleware.Logger("admin"))
		admin.Use(middleware.Cors())
		admin.GET("/login/index", login.Get)
		admin.PUT("/login/index", login.Put)

		admin.GET("/nav/index", nav.Get)
		admin.POST("/nav/index", nav.Post)
		admin.PUT("/nav/index", nav.Put)
	}
}
