package connection

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)
func ConfigCors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	})
}
