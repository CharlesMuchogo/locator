package routes

import (
	"github.com/gin-gonic/gin"
	"main.go/controllers"
	"main.go/location"
	"main.go/middlewares"
	"main.go/user"
)

func InitRouter() *gin.Engine {

	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		api.POST("/register", user.Signup)
		api.POST("/login", user.Login)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.POST("/location", location.UpdateLocation)
			secured.GET("/location", location.GetLocation)
			secured.POST("/user", user.UpdateProfile)
			secured.POST("/request_promotion", user.RequestPromotion)
			secured.GET("/promote_user", user.PromoteUser)
		}
	}

	return router
}
