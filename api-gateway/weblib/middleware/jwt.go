package middleware

import (
	"api-gateway/pkg/utils"

	"github.com/gin-gonic/gin"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code uint32

		code = 200
		//Get token
		token := c.GetHeader("Authorization")

		//Not Header:Authorization
		if token == "" {
			code = 404
		} else {
			//ParseToken,if token is wrong
			_, err := utils.ParseToken(token)
			if err != nil {
				code = 401
			}
		}

		//code
		if code != 200 {
			c.JSON(500, gin.H{
				"code": code,
				"msg":  "鉴权失败",
			})

			c.Abort()
			return
		}

		//token existed and successfully parsed
		c.Next()
	}
}
