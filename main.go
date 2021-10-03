package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func QueryMarket(c *gin.Context) {
	fmt.Println("QueryMarket router")
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {
	fmt.Println("Hello market-tracker project")
	app := gin.Default()
	app.GET("/query-market", QueryMarket)
	app.Run(":3000")
}
