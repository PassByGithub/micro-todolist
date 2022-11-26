package core

import (
	"context"
	"errors"
	"user/model"
	"user/services"

	"gorm.io/gorm"
)

//序列化
func BuildUser(item model.User) *services.UserModel {
	userModel := services.UserModel{
		ID:        uint32(item.ID),
		Username:  item.UserName,
		CreateAt:  item.CreatedAt.Unix(),
		UpdatedAt: item.UpdatedAt.Unix(),
	}

	return &userModel
}

func (*UserService) UserLogin(ctx context.Context, req *services.UserRequest, res *services.UserResponse) error {
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

	if user.CheckPassword(req.UserPassword) == false {
		res.Code = 400
		return nil
	}

	res.UserDetail = BuildUser(user)
	return nil

}

func (*UserService) UserRegister(ctx context.Context, req *services.UserRequest, res *services.UserResponse) error {

	var user model.User

	if req.UserPassword != req.PasswordConfirm {
		err := errors.New(("两次输入密码不一致"))
		return err
	}

	result := model.DB.Model(&model.User{}).Where("user_name=?", req.UserName)
	count := result.RowsAffected

	if err := result.Error; err != nil {
		return err
	}

	if count > 0 {
		err := errors.New("用户名已存在")
		return err
	}
	if err := user.SetPassword(req.UserPassword); err != nil {
		return err
	}
	if err := model.DB.Create(&user).Error; err != nil {
		return err
	}
	res.UserDetail = BuildUser(user)
	return nil

}
