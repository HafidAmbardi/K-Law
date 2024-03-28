package routes

import (
	"OTI-inbound/controller"
	middlewares "OTI-inbound/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.Use(func(ctx *gin.Context) {
		ctx.Set("db", db)
	})

	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	r.PUT("/:id", controller.UpdatePassword)

	r.POST("/registeradmin", controller.RegisterAdmin)
	r.POST("/loginadmin", controller.LoginAdmin)
	updatePassMiddlewareRoute := r.Group("update-password-admin")
	updatePassMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	updatePassMiddlewareRoute.POST("/:id", controller.UpdatePassAdmin)

	r.GET("/categories", controller.GetAllCategories)
	categoriesMiddlewareRoute := r.Group("/categories")
	categoriesMiddlewareRoute.DELETE("/:id", controller.DeleteCategories)
	categoriesMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	categoriesMiddlewareRoute.POST("/", controller.CreateCategories)
	categoriesMiddlewareRoute.PUT("/:id", controller.UpdateCategories)

	r.GET("/post", controller.GetAllPost)
	r.GET("/post/:text", controller.GetPostByText)
	r.POST("/post", controller.CreatePost)
	r.PUT("/post/:id", controller.UpdatePost)
	r.DELETE("/post/:id", controller.DeletePost)

	r.GET("/comment", controller.GetAllComment)
	r.POST("/comment", controller.CreateComment)
	r.PUT("/comment/:id", controller.UpdateComment)
	r.DELETE("/comment/:id", controller.DeleteComment)

	r.GET("/vote", controller.GetAllVote)
	r.GET("/countvote", controller.CountVote)
	r.POST("/vote", controller.CreateVote)
	r.DELETE("/vote/:id", controller.DeleteVote)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r

}
