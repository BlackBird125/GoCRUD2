package server

import (
	controllers "github.com/BlackBird125/GoCRUD2/controller"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init() {
	r := router()
	r.Run(":8080")
}

func router() *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	u := r.Group("/users")
	{
		ctrl := controllers.UserController{}
		u.GET("", ctrl.Index)
		u.GET("/:id", ctrl.Show)
		u.POST("", ctrl.Create)
		u.PUT("/:id", ctrl.Update)
		u.DELETE("/:id", ctrl.Delete)
	}

	// p := r.Group("/posts")
	// {
	// 	ctrl := controllers.PostController{}
	// 	p.GET("", ctrl.Index)
	// 	p.POST("", ctrl.Create)
	// 	p.GET("/:id", ctrl.Show)
	// 	p.PUT("/:id", ctrl.Update)
	// 	p.DELETE("/:id", ctrl.Delete)
	// }

	return r
}