package routers

import (
	"nowqs/frame/config"
	admin "nowqs/frame/http/admin/routers"
	swagger_ui "nowqs/frame/http/assets"
	blog "nowqs/frame/http/blog/routers"
	"nowqs/frame/http/middleware"
	ws "nowqs/frame/http/websocket/routers"

	assetfs "github.com/elazarl/go-bindata-assetfs"

	"github.com/gin-gonic/gin"
)

// NewRouter is router function
func NewRouter(r *gin.Engine) *gin.Engine {
	r.Use(middleware.Cors())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// admin routers
	admin.NewRouters(r)
	blog.NewRouters(r)
	ws.NewRouters(r)

	swaggerUi := assetfs.AssetFS{Asset: swagger_ui.Asset, AssetDir: swagger_ui.AssetDir, AssetInfo: swagger_ui.AssetInfo, Prefix: "http/assets/swagger-ui"}

	// swagger router group
	swagger := r.Group("/swagger")
	{
		adminPath := config.GetHttpPath() + config.PathSeparator + "admin"
		blogV1Path := config.GetHttpPath() + config.PathSeparator + "blog" + config.PathSeparator + "v1"
		swagger.StaticFS("/ui", &swaggerUi)
		// swagger.StaticFS("/ui", http.Dir(config.GetHttpPath()+config.PathSeparator+"assets"+config.PathSeparator+"swagger-ui"))
		swagger.StaticFile("/admin.json", adminPath+config.PathSeparator+"swagger"+config.PathSeparator+"swagger.json")
		swagger.StaticFile("/blog/v1.json", blogV1Path+config.PathSeparator+"swagger"+config.PathSeparator+"swagger.json")
	}
	return r
}
