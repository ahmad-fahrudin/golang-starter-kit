package routes

import (
	"golang-starter-kit/src/controller"
	"golang-starter-kit/src/middleware"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRoutes configures all application routes
func SetupRoutes(
	router *gin.Engine,
	userController *controller.UserController,
	authController *controller.AuthController,
	jwtSecret string,
) {
	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Server is running",
		})
	})

	// Swagger endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API version 1
	v1 := router.Group("/api/v1")
	{
		// Authentication routes (public)
		auth := v1.Group("/auth")
		{
			auth.POST("/register", authController.Register)
			auth.POST("/login", middleware.RateLimitMiddleware(5, 15*time.Minute), authController.Login)
			auth.POST("/logout", authController.Logout)
		}

		// User routes (public)
		users := v1.Group("/users")
		{
			users.POST("", userController.CreateUser)                        // Create user
			users.POST("/pagination", userController.GetUsersWithPagination) // Get users with pagination
			users.GET("/:id", userController.GetUser)                        // Get user by ID
			users.PUT("/:id", userController.UpdateUser)                     // Update user
			users.DELETE("/:id", userController.DeleteUser)                  // Delete user
		}

		// Protected routes
		protected := v1.Group("")
		protected.Use(middleware.AuthMiddleware(jwtSecret))
		{
			// Profile routes (protected)
			protected.GET("/profile", userController.GetProfile)
			protected.PUT("/profile", userController.UpdateProfile)
		}
	}
}
