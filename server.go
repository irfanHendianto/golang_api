package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golang_api/config"
	"github.com/golang_api/controllers"
	"github.com/golang_api/middleware"
	"github.com/golang_api/repository"
	"github.com/golang_api/service"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                   = config.SetUpDatabaseConnection()
	userRepository repository.UserRepository  = repository.NewUserRepository(db)
	jwtService     service.JWTService         = service.NewJWTService()
	authService    service.AuthService        = service.NewAuthService(userRepository)
	authController controllers.AuthController = controllers.NewAuthController(authService, jwtService)
	userService    service.UserService        = service.NewUserService(userRepository)
	userController controllers.UserController = controllers.NewUserController(userService, jwtService)
)

func main() {
	r := gin.Default()
	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
		authRoutes.DELETE("/delete", userController.Delete)
	}

	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
