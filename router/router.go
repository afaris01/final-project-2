package router

import (
	"final-project-2/controllers"
	"final-project-2/middlewares"

	"github.com/gin-gonic/gin"
)

func MulaiApp() *gin.Engine {
	router := gin.Default()

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
		userRouter.PUT("/:user_id", middlewares.Authentication(), controllers.UbahUser)
		userRouter.DELETE("/:user_id", middlewares.Authentication(), controllers.HapusUser)
	}

	photoRouter := router.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.POST("/", controllers.BuatPhoto)
		photoRouter.GET("/", controllers.AmbilPhoto)
		photoRouter.PUT("/:photo_id", middlewares.PhotoAuthorization(), controllers.UbahPhoto)
		photoRouter.DELETE("/:photo_id", middlewares.PhotoAuthorization(), controllers.HapusPhoto)
	}

	commentRouter := router.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.POST("/", controllers.BuatComment)
		commentRouter.GET("/", controllers.AmbilComment)
		commentRouter.PUT("/:comment_id", middlewares.CommentAuthorization(), controllers.UbahComment)
		commentRouter.DELETE("/:comment_id", middlewares.CommentAuthorization(), controllers.HapusComment)
	}

	socialMediaRouter := router.Group("/socialmedias")
	{
		socialMediaRouter.Use(middlewares.Authentication())
		socialMediaRouter.POST("/", controllers.BuatSocialMedia)
		socialMediaRouter.GET("/", controllers.AmbilSocialMedia)
		socialMediaRouter.PUT("/:socialmedia_id", middlewares.SocialMediaAuthorization(), controllers.UbahSocialMedia)
		socialMediaRouter.DELETE("/:socialmedia_id", middlewares.SocialMediaAuthorization(), controllers.HapusSocialMedia)
	}
	return router
}
