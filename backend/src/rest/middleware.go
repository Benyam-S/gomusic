package rest

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// MiddleWareLogger is just a simple gin middleware
func MiddleWareLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Getting inside")
		c.Next()
		fmt.Println("Moving outside")
	}
}

// CORSMiddleware is a cosr middleware
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers",
			"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, access-control-allow-origin, access-control-allow-headers")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
