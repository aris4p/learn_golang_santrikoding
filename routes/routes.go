package routes

import (
	"github.com/aris4p/controllers"
	"github.com/aris4p/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// initialize gin
	router := gin.Default()

	// route register
	router.POST("/api/register", controllers.Register)
	// route login
	router.POST("/api/login", controllers.Login)

	// route users
	router.GET("/api/users", middlewares.AuthMiddleware(), controllers.FindUser)

	// route user create
	router.POST("/api/users", middlewares.AuthMiddleware(), controllers.CreateUser)

	// route user by id
	router.GET("/api/users/:id", middlewares.AuthMiddleware(), controllers.FindUserById)

	// route user update
	router.PUT("/api/users/:id", middlewares.AuthMiddleware(), controllers.UpdateUser)

	// route user delete
	router.DELETE("/api/users/:id", middlewares.AuthMiddleware(), controllers.DeleteUser)

	return router
}
