package routes

import (
	"Simple_Web_API_Login/controllers"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRoutes := r.Group("/auth")
	{
		userRoutes.POST("/login", controllers.Login)
		userRoutes.POST("/register", controllers.Register)
	}

	return r
}
