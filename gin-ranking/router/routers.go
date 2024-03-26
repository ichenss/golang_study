package router

import (
	"gin-ranking/controllers"
	"gin-ranking/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.Use(gin.LoggerWithConfig(logger.LoggerToFile()))
	r.Use(logger.Recover)

	r.GET("/url", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello gin")
	})

	// 分组
	user := r.Group("/user")
	{
		user.GET("/info/:id/:name", controllers.User{}.GetUserInfo)

		user.POST("/list", controllers.User{}.GetList)

		user.PUT("/add", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "user add")
		})

		user.DELETE("/delete", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "user delete")
		})
	}

	od := r.Group("/order")
	{
		od.POST("/list", controllers.OrderController{}.GetList)
	}

	return r
}
