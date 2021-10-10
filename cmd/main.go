package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(static.Serve("/", static.LocalFile("dist", false)))

	r.GET("/api/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "hello",
		})
	})

	if err := r.Run(":3000"); err != nil {
		panic(err)
	}
}
