package weblib

import (
	"api-gateway/weblib/handlers"
	"api-gateway/weblib/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter(service ...interface{}) *gin.Engine {

	//gin接口配置调用中间件
	ginRouter := gin.Default()
	ginRouter.Use(middleware.Cors(), middleware.InitMiddleware(service), middleware.ErrorMiddleware())
	store := cookie.NewStore([]byte("secret"))
	ginRouter.Use(sessions.Sessions("mysession", store))

	v1 := ginRouter.Group("/api/v1")
	{
		v1.GET("ping", func(context *gin.Context) {
			context.JSON(200, "success")
		})

		//user service
		v1.POST("/user/register", handlers.UserRegister)
		v1.POST("/user/login", handlers.UserLogin)

		//login authentification
		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			authed.GET("task", handlers.GetTaskList)
			authed.POST("task", handlers.CreateTaskList)
			authed.GET("task/:id", handlers.GetTaskDetail)
			authed.PUT("task/:id", handlers.UpdateTask)
			authed.DELETE("task/:id", handlers.DeleteTask)

		}
	}

	return ginRouter

}
