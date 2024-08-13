package main

import (
	"github.com/gin-gonic/gin"
	"github.com/MartoneWorkshop/jwt-go-test/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	
	r.GET("/ping", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "pong",
		})		
	})
	r.Run()
}