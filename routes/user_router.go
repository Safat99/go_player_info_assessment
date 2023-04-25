package routes

import "github.com/gin-gonic/gin"

func SetupUserRouter(userService *service.UserService) *gin.Engine {
	r := gin.Default()

	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/register", userService.Create)
		userRoutes.POST("/login", userService.Login)
	}

	return r
}
