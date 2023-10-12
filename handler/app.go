package handler

import (
	"finalProject2/database"
	middlewares "finalProject2/pkg/middleware"
	"finalProject2/repository/photo_repository/photo_pg"
	"finalProject2/repository/user_repository/user_pg"
	"finalProject2/service"

	"github.com/gin-gonic/gin"
)

func StartApp() {
	database.InitiliazeDatabase()
	db := database.GetDatabaseInstance()

	// user
	userRepo := user_pg.NewOrderPG(db)
	userService := service.NewUserService(userRepo)
	userHandler := NewUserHandler(userService)

	// photo
	photoRepo := photo_pg.NewPhotoPG(db)
	photoService := service.NewPhotoService(photoRepo)
	photoHandler := NewPhotoHandler(photoService)

	r := gin.Default()

	//user
	r.POST("/users/register", userHandler.register)
	r.POST("/users/login", userHandler.Login)

	users := r.Group("/users")
	photos := r.Group("/photos")
	users.Use(middlewares.Authentication())
	{
		users.PUT("", userHandler.EditUser)
		users.DELETE("", userHandler.DeleteUser)
	}

	photos.Use(middlewares.Authentication())
	{
		photos.POST("", photoHandler.CreatePhoto)
		photos.GET("", photoHandler.GetPhotos)
		photos.PUT("/:photoId", photoHandler.EditPhoto)
		photos.DELETE("/:photoId", photoHandler.DeletePhoto)
	}

	r.Run(":8080")

}
