package routers

import (
	adminController "nowqs/frame/http/admin/controller"
	"nowqs/frame/http/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouters is admin routers function
func NewRouters(r *gin.Engine) *gin.Engine {

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	login := adminController.NewLogin()
	nav := adminController.NewNav()
	authGroup := adminController.NewAuthGroup()
	admin := r.Group("/admin")
	{
		admin.Use(middleware.Logger("admin"))
		admin.Use(middleware.Cors())
		admin.GET("/login/index", login.Get)
		admin.PUT("/login/index", login.Put)

		admin.GET("/nav/index", nav.Get)
		admin.POST("/nav/index", nav.Post)
		admin.PUT("/nav/index", nav.Put)
		admin.DELETE("/nav/index", nav.Delete)

		admin.GET("/auth_group/index", authGroup.Get)
		admin.POST("/auth_group/index", authGroup.Post)
		admin.PUT("/auth_group/index", authGroup.Put)
		admin.DELETE("/auth_group/index", authGroup.Delete)

	}

	return r
}
