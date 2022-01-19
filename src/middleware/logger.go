package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Set("example", "123456")

		c.Next()

		latency := time.Since(t)
		log.Print("Result in :", latency)

		status := c.Writer.Status()
		log.Print(status)
	}
}
