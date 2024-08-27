package router

import (
	"loan_tracker_api/deliveries/controllers"
	"loan_tracker_api/infrastructure"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetRouter(router *gin.Engine, cu *controllers.UserController, client *mongo.Client) {

	router.POST("/user/register", cu.RegisterUser)
	router.POST("/user/verify-email", cu.VerifyEmail)
	router.POST("/user/login", cu.LoginUser)
	router.GET("/user/token-refresh", cu.TokenRefresh)
	router.GET("/user/profile", infrastructure.AuthMiddleware(client), cu.UserProfile)
	router.GET("/user/logout", infrastructure.AuthMiddleware(client), cu.LogoutUser)
	router.PUT("/user/update", infrastructure.AuthMiddleware(client), cu.UpdateUserDetails)

	router.POST("/user/password-reset", cu.ForgotPassword)
	router.POST("/user/password-update", cu.ResetPassword)

	admino := router.Group("/admin")
	admino.Use(infrastructure.AuthMiddleware(client), infrastructure.AdminMiddleware)
	{
		admino.GET("/users", cu.ViewAllUsers)
		admino.DELETE("/user/:id", cu.DeleteUser)
	}

}
