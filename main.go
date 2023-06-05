package main

import (
	"simple-payment/delivery"
)

func main() {
	// r := gin.Default()
	// r.GET("", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "Hello World",
	// 	})
	// })
	// r.Run("localhost:8080")
	delivery.NewServer("localhost:8080").Run()
}
