package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/irfan44/task-5-vix-btpns-IrfanNurghiffariM/controllers"
	"github.com/irfan44/task-5-vix-btpns-IrfanNurghiffariM/middlewares"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	router.POST("/users/login", controllers.Login)
	router.POST("/users/register", controllers.CreateUser)
	router.PUT("/users/:userId", controllers.UpdateUser)
	router.DELETE("/users/:userId", controllers.DeleteUser)

	router.GET("/photos", controllers.GetPhoto)

	secured := router.Group("/").Use(middlewares.AuthMiddleware())
	{
		secured.POST("/photos", controllers.CreatePhoto)
		secured.PUT("/photos/:photoId", controllers.UpdatePhoto)
		secured.DELETE("/photos/:photoId", controllers.DeletePhoto)
	}

	return router
}
