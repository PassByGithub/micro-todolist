package handlers

import (
	"api-gateway/pkg/utils"
	"context"
	"net/http"
	"proto/microuser"

	"github.com/gin-gonic/gin"
)

func UserRegister(ctx *gin.Context) {
	var userReq microuser.UserRequest
	PanicIfUserError(ctx.Bind(&userReq))

	//从gin.Keys中取出请求的服务实例
	userService := ctx.Keys["userService"].(microuser.UserService)
	userResp, err := userService.UserRegister(context.Background(), &userReq)

	//return resp
	PanicIfUserError(err)
	ctx.JSON(http.StatusOK, gin.H{"data": userResp})
}

func UserLogin(ctx *gin.Context) {
	var userReq microuser.UserRequest
	PanicIfUserError(ctx.Bind(&userReq))
	userService := ctx.Keys["userService"].(microuser.UserService)
	userResp, err := userService.UserLogin(context.Background(), &userReq)
	PanicIfUserError(err)

	//生成token
	token, err := utils.GenerateToken(uint(userResp.UserDetail.ID))
	if err != nil {
		panic(err)
	}

	//返回token
	ctx.JSON(http.StatusOK, gin.H{
		"code": userResp.Code,
		"msg":  "token generated",
		"data": gin.H{
			"user":  userResp.UserDetail,
			"token": token,
		},
	})
}
