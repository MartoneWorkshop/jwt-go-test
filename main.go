package main

import (
	"net/http"
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
	
	r.POST("/signup", controllers.Signup)
	
	r.Run()
}