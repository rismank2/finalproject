package router

import (
	controllers "finalproject/controller"
	"finalproject/middleware"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
		userRouter.PUT("/:userId", middleware.Authentication(), controllers.UpdateUser)
		userRouter.DELETE("/:userId", middleware.Authentication(), controllers.DeleteUser)
	}
	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middleware.Authentication())
		photoRouter.POST("/", controllers.CreatePhoto)
		photoRouter.GET("/", controllers.GetPhotos)
		photoRouter.PUT("/:photoId", middleware.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middleware.PhotoAuthorization(), controllers.DeletePhoto)
	}

	commentRouter := r.Group("/comment")
	{
		commentRouter.Use(middleware.Authentication())
		commentRouter.POST("/", controllers.CreateComment)
		commentRouter.GET("/", controllers.GetComments)
		commentRouter.PUT("/:commentId", middleware.CommentAuthorization(), controllers.UpdateComment)
		commentRouter.DELETE("/:commentId", middleware.CommentAuthorization(), controllers.DeleteComment)
	}

	socialMediaRouter := r.Group("/socialmedias")
	{
		socialMediaRouter.Use(middleware.Authentication())
		socialMediaRouter.POST("/", controllers.CreateSocialMedia)
		socialMediaRouter.GET("/", controllers.GetSocialMedia)
		socialMediaRouter.PUT("/:socialMediaId", middleware.SocialMediaAuthorization(), controllers.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:socialMediaId", middleware.SocialMediaAuthorization(), controllers.DeleteSocialMedia)
	}
	return r
}
