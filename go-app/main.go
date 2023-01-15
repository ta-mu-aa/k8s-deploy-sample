package main

import (
	"log"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("start server...")
	r := gin.Default()
	r.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	log.Fatal(r.Run())
}
