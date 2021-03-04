package http

import (
	"nowqs/frame/config"
	"nowqs/frame/http/routers"
	"nowqs/frame/models/mysql"
	"nowqs/frame/models/redis"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Run is http run function
func Run() {
	// initialization mysql
	mysql.InitDb()

	// initialization redis
	redis.InitPool()

	r := gin.Default()
	r = routers.NewRouter(r)

	r.Run(config.AppConfig.Host + ":" + strconv.Itoa(config.AppConfig.Port))
}
