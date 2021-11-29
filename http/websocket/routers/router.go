package routers

import (
	v1 "nowqs/frame/http/websocket/v1/controller"

	"github.com/gin-gonic/gin"
)

// NewRouters is websocket routers function
func NewRouters(r *gin.Engine) {
	v1Hub := v1.NewHub()
	go v1Hub.Run()

	r.GET("wss", func(c *gin.Context) {
		v1.WsHandle(v1Hub, c)
	})
}
