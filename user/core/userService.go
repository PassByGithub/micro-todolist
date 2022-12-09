package core

import (
	"context"
	"errors"
	"fmt"
	"user/model"
	"user/service"

	"gorm.io/gorm"
)

//序列化
//transform model.User into service.UserModel to microservice
func BuildUser(item model.User) *service.UserModel {
	userModel := service.UserModel{
		ID:        uint32(item.ID),
		Username:  item.UserName,
		CreateAt:  item.CreatedAt.Unix(),
		UpdatedAt: item.UpdatedAt.Unix(),
	}

	return &userModel
}

func (*UserService) UserLogin(ctx context.Context, req *service.UserRequest, res *service.UserResponse) error {
	var user model.User
	res.Code = 200

	if err := model.DB.Where("user_name=?", req.UserName).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			res.Code = 400
			return nil
		}
		res.Code = 500
		return nil
	}

	if !user.CheckPassword(req.UserPassword) {
		res.Code = 400
		return nil
	}

	res.UserDetail = BuildUser(user)
	return nil

}

//req != model.User
//req recieved by microservice
//model.User是数据库模型
func (*UserService) UserRegister(ctx context.Context, req *service.UserRequest, res *service.UserResponse) error {

	var user model.User

	//Confirm password
	if req.UserPassword != req.PasswordConfirm {
		err := errors.New(("两次输入密码不一致"))
		return err
	}

	//Find user if existed
	//model.DB = db
	//table = user
	var count int64
	result := model.DB.Model(&model.User{}).Where("user_name=?", req.UserName).Count(&count)

	if err := result.Error; err != nil {
		return err
	}

	if count > 0 {
		fmt.Println(count)
		err := errors.New("用户名已存在")
		return err
	}
	fmt.Println(count)

	//Password Encrypted
	if err := user.SetPassword(req.UserPassword); err != nil {
		return err
	} else {
		user.UserName = req.UserName
	}

	//User created into database
	if err := model.DB.Create(&user).Error; err != nil {
		return err
	}

	//Flush res and return nil
	res.UserDetail = BuildUser(user)
	return nil

}
