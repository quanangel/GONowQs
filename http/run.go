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
	// gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	// 	fmt.Printf("endpoint %v %v %v %v'\r\n", httpMethod, absolutePath, handlerName, nuHandlers)
	// }
	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)
	r := gin.Default()
	r = routers.NewRouter(r)

	r.Run(config.AppConfig.Host + ":" + strconv.Itoa(config.AppConfig.Port))
}
