package routes

import (
	"github.com/Rashad-Muntar/my-go-rest/controllers"
	"github.com/Rashad-Muntar/my-go-rest/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine){
	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)
	router.GET("/users", controllers.GetUsers)
	router.PUT("/users/:id", controllers.UpdateUser)
	router.DELETE("/users/:id", controllers.DeleteUser)
	router.GET("/validate", middleware.RequireAuth, controllers.ValidateAuth)
}