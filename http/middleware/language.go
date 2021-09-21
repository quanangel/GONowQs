package middleware

import "github.com/gin-gonic/gin"

// Language is language middle
func Language() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: change language
		c.Next()
	}
}
