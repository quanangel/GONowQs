package middleware

import (
	"github.com/gin-gonic/gin"
)

// Logger is log function
func Logger() gin.HandlerFunc {
	// logPath := config.GetLogPath()
	// TODO: not finish
	return func(c *gin.Context) {
		// now := time.Now()
		// logDir := logPath + config.PathSeparator + now.Format("200601") + config.PathSeparator + "admin"
		// err := os.Mkdir(logDir, 0666)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// logName := logDir + config.PathSeparator + strconv.Itoa(now.Day()) + ".log"
		// logFile, err := os.OpenFile(logName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		// defer logFile.Close()
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// gin.DefaultWriter = io.MultiWriter(logFile)
		c.Next()
	}
}
