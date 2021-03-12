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
	// TODO: not finish

	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)
	// r := gin.Default()
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r = routers.NewRouter(r)

	r.Run(config.AppConfig.Host + ":" + strconv.Itoa(config.AppConfig.Port))
}
