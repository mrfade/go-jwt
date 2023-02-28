package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mrfade/go-jwt/controllers"
	"github.com/mrfade/go-jwt/initializers"
	"github.com/mrfade/go-jwt/middlewares"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middlewares.RequireAuth, controllers.Validate)

	r.Run()
}
