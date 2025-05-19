package routes

import (
	"github.com/aris4p/controllers"
	"github.com/aris4p/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// initialize gin
	router := gin.Default()

	// set up CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	// Auth routes
	router.POST("/api/register", controllers.Register)
	router.POST("/api/login", controllers.Login)

	// Serve folder uploads
	router.Static("/uploads", "./uploads")

	// User routes (with authentication)
	userRoutes := router.Group("/api/users", middlewares.AuthMiddleware())
	{
		userRoutes.GET("", controllers.FindUser)
		userRoutes.POST("", controllers.CreateUser)
		userRoutes.GET("/:id", controllers.FindUserById)
		userRoutes.PUT("/:id", controllers.UpdateUser)
		userRoutes.DELETE("/:id", controllers.DeleteUser)
	}

	// Product routes (with authentication)
	productRoutes := router.Group("/api/products", middlewares.AuthMiddleware())
	{
		productRoutes.GET("", controllers.FindProduct)
		productRoutes.POST("", controllers.CreateProduct)
		productRoutes.GET("/:id", controllers.FindProductById)
		productRoutes.PUT("/:id", controllers.UpdateProduct)
		productRoutes.DELETE("/:id", controllers.DeleteProduct)
	}

	return router
}
