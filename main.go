package main

import (
  	"github.com/gin-gonic/gin"
	"task-5-vix-btpns-andrian/controllers"
	"task-5-vix-btpns-andrian/models"
	"task-5-vix-btpns-andrian/middlewares"
)

func main() {

	models.ConnectDataBase()
	
	r := gin.Default()

	public := r.Group("/users")

	public.POST("/register", controllers.Register)
	public.POST("/login",controllers.Login)
	// jwt
	public.Use(middlewares.JwtAuthMiddleware())
	public.GET("/profile",controllers.CurrentUser)
	
	r.Run(":8080")

}