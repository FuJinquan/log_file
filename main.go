package main

import (
	"io"
	"os"

	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.DisableConsoleColor()
	file, err := os.Create("gin.log")
	if err != nil {
		fmt.Println("Failed to create a file:%v", err)
	}
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)

	r := gin.Default()
	r.GET("/log", func(c *gin.Context) {
		c.String(200, "successful!")
	})
	r.Run(":8080")
}
