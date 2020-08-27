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
