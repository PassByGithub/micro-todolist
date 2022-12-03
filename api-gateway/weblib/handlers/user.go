package handlers

import (
	"api-gateway/pkg/utils"
	"api-gateway/services"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRegister(ctx *gin.Context) {
	var userReq services.UserRequest
	PanicIfUserError(ctx.Bind(&userReq))

	//从gin.Keys中取出请求的服务实例
	userService := ctx.Keys["userService"].(services.UserService)
	userResp, err := userService.UserRegister(context.Background(), &userReq)
	PanicIfUserError(err)
	ctx.JSON(http.StatusOK, gin.H{"data": userResp})
}

func UserLogin(ctx *gin.Context) {
	var userReq services.UserRequest
	PanicIfUserError(ctx.Bind(&userReq))
	userService := ctx.Keys["userService"].(services.UserService)
	userResp, err := userService.UserLogin(context.Background(), &userReq)
	PanicIfUserError(err)

	token, err := utils.GenerateToken(uint(userResp.UserDetail.ID))

	ctx.JSON(http.StatusOK, gin.H{
		"code": userResp.Code,
		"msg":  "token generated",
		"data": gin.H{
			"user":  userResp.UserDetail,
			"token": token,
		},
	})
}