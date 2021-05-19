package routers

import (
	"net/http"
	"nowqs/frame/config"
	admin "nowqs/frame/http/admin/routers"
	blog "nowqs/frame/http/blog/routers"

	"github.com/gin-gonic/gin"
)

// NewRouter is router function
func NewRouter(r *gin.Engine) *gin.Engine {

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// admin routers
	r = admin.NewRouters(r)
	r = blog.NewRouters(r)

	// swagger router group
	swagger := r.Group("/swagger")
	{
		adminPath := config.GetHttpPath() + config.PathSeparator + "admin"
		blogV1Path := config.GetHttpPath() + config.PathSeparator + "blog" + config.PathSeparator + "v1"
		swagger.StaticFS("/ui", http.Dir(config.GetHttpPath()+config.PathSeparator+"assets"+config.PathSeparator+"swagger-ui"))
		swagger.StaticFile("/admin.json", adminPath+config.PathSeparator+"swagger"+config.PathSeparator+"swagger.json")
		swagger.StaticFile("/blog/v1.json", blogV1Path+config.PathSeparator+"swagger"+config.PathSeparator+"swagger.json")
	}
	return r
}
