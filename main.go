package main

import (
	"github.com/gin-gonic/gin"
	"project2/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	
	r.POST("/signup", controllers.Signup)
	
	r.Run()
}