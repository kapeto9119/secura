package middlewares

import (
	"log"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

// Recover returns a middleware that recovers from panics
func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Log the error and stack trace
				debugStack := debug.Stack()
				log.Printf("Recovery from panic: %v\nStack trace: %s", err, debugStack)

				// Return a 500 error
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": "Internal Server Error",
				})
			}
		}()
		c.Next()
	}
}
