package routers

import (
	admin "nowqs/frame/http/admin/routers"

	"github.com/gin-gonic/gin"
)

// NewRouter is router function
func NewRouter(r *gin.Engine) *gin.Engine {
	r.StaticFile("/swagger", "./docs/swagger.json")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r = admin.NewRouters(r)
	return r
}
