package middleware

import (
	"fmt"
	"nowqs/frame/config"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger is log function
func Logger() gin.HandlerFunc {
	logPath := config.GetLogPath()
	return func(c *gin.Context) {
		// logBody := make(map[string]interface{})

		now := time.Now()
		logDir := logPath + config.PathSeparator + now.Format("200601") + config.PathSeparator + "admin"
		err := os.Mkdir(logDir, 0666)
		if err != nil {
			fmt.Println(err)
		}
		logName := logDir + config.PathSeparator + strconv.Itoa(now.Day()) + ".log"
		logFile, err := os.OpenFile(logName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		defer logFile.Close()
		if err != nil {
			fmt.Fprintln(logFile, err)
		}
		c.Next()

		endTime := time.Now()
		// logBody := map[string]interface{}{
		// 	"request_time":      now.Format("2006-01-02 15:04:05"),
		// 	"request_method":    c.Request.Method,
		// 	"request_uri":       c.Request.RequestURI,
		// 	"request_client_ip": c.ClientIP(),
		// 	"request_proto":     c.Request.Proto,
		// 	"request_post":      c.Request.PostForm.Encode(),
		// 	"response_time":     endTime.Format("2006-01-02 15:04:05"),
		// 	"response_code":     c.Writer.Status(),
		// 	"take_time":         fmt.Sprintf("%13v", endTime.Sub(now)),
		// }
		// result, err := json.Marshal(logBody)
		// if err != nil {
		// 	fmt.Fprintln(logFile, err)
		// }
		// jsonStringData := string(result)
		// fmt.Fprintln(logFile, jsonStringData)
		logContent := fmt.Sprintf("[%s-%s] %s | %v | %13v | %s | %s | %s", c.Request.Proto, c.Request.Method, now.Format("2006-01-02 15:04:05"), c.Writer.Status(), endTime.Sub(now), c.ClientIP(), c.Request.RequestURI, c.Request.PostForm.Encode())
		fmt.Fprintln(logFile, logContent)
	}
}
