package router

import (
	"final-project-2/controllers"
	"final-project-2/middlewares"

	"github.com/gin-gonic/gin"
)

func MulaiApp() *gin.Engine{
	router := gin.Default()

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
		userRouter.PUT("/:userId", middlewares.Authentication(), controllers.UbahUser)
		userRouter.DELETE("/:userId", middlewares.Authentication(), controllers.HapusUser)
	}

	photoRouter := router.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.POST("/", controllers.BuatPhoto)
		photoRouter.GET("/", controllers.AmbilPhoto)
		photoRouter.PUT("/:photoId", middlewares.PhotoAutorization(), controllers.UbahPhoto)
		photoRouter.DELETE("/:photoId", middlewares.PhotoAutorization(), controllers.HapusPhoto)
	}

	commentRouter := router.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.POST("/", controllers.BuatComment)
		commentRouter.GET("/", controllers.AmbilComment)
		commentRouter.PUT("/:commentId", middlewares.CommentAuthorization(), controllers.UbahComment)
		commentRouter.DELETE("/:commentId", middlewares.CommentAuthorization(), controllers.HapusComment)
	}

	socialMediaRouter := router.Group("/socialmedias")
	{
		socialMediaRouter.Use(middlewares.Authentication())
		socialMediaRouter.POST("/", controllers.BuatSocialMedia)
		socialMediaRouter.GET("/", controllers.AmbilSocialMedia)
		socialMediaRouter.PUT("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.UbahSocialMedia)
		socialMediaRouter.DELETE("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.HapusSocialMedia)
	}
	return router
}