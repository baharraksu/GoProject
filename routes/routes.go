package routes

import (
	"GoProject/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Rotaları tanımlama
func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	r.POST("/users", func(c *gin.Context) {
		controllers.CreateUser(c, db)
	})
	r.GET("/users", func(c *gin.Context) {
		controllers.GetUsers(c, db)
	})
	r.GET("/users/:id", func(c *gin.Context) {
		controllers.GetUserByID(c, db)
	})
	r.PUT("/users/:id", func(c *gin.Context) {
		controllers.UpdateUser(c, db)
	})
	r.DELETE("/users/:id", func(c *gin.Context) {
		controllers.DeleteUser(c, db)
	})
}
