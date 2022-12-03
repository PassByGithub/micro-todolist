package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

//传入一个空接口
func InitMiddleware(service []interface{}) gin.HandlerFunc {
	return func(context *gin.Context) {
		//Pass instance into the gin.Keys
		context.Keys = make(map[string]interface{})
		context.Keys["userService"] = service[0]
		context.Next()
	}
}

//Middleware for error handling
func ErrorMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				context.JSON(200, gin.H{
					"code": 404,
					"msg":  fmt.Sprintf("%s", err),
				})
				context.Abort()
			}
		}()
		context.Next()
	}

}
