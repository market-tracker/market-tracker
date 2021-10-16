package health

import (
	"github.com/gin-gonic/gin"
)

func Handler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "server is up and running",
	})
}
