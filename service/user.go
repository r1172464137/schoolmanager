package service

import (
	"school_manager/model"
	"school_manager/pkg"
	"school_manager/serializer"
	"school_manager/util"

	"gorm.io/gorm"
)

type UserService struct {
	Uid      uint   `json:"uid" form:"uid" binding:""`
	Name     string `json:"name" form:"name" binding:""`
	UserName string `json:"username" form:"username" binding:"required"`
	Capacity bool   `json:"identity" form:"identity" binding:""`
	Password string `json:"password" form:"password" binding:"required"`
	College  string `json:"college" form:"college" binding:""`
}

func (service *UserService) Register() *serializer.Response {
	code := pkg.SUCCESS
	var user model.User
	var count int64
	model.DB.Model(&model.User{}).Where("username = ?", service.UserName).First(&user).Count(&count)
	if count == 1 {
		code = pkg.ErrorExistUser
		return &serializer.Response{
			Status: code,
			Msg:    pkg.GetMsg(code),
		}
	} //查看是否存在同名用户
	model.DB.Model(&model.User{}).Where("uid = ?", service.Uid).First(&user).Count(&count)
	if count == 1 {
		code = pkg.ErrorExistUser
		return &serializer.Response{
			Status: code,
			Msg:    pkg.GetMsg(code),
		}
	} //查看是否存在同样的id
	user.Username = service.UserName
	user.Uid = service.Uid
	user.Name = service.Name
	user.Capacity = service.Capacity
	user.College = service.College
	if err := user.SetPassword(service.Password); err != nil {
		util.LogrusObj.Info(err)
		code = pkg.ErrorFailEncryption
		return &serializer.Response{
			Status: code,
			Msg:    pkg.GetMsg(code),
		}
	} //不对密码进行明文存储，进行加密操作

	if err := model.DB.Create(&user).Error; err != nil {
		util.LogrusObj.Info(err)
		code = pkg.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg:    pkg.GetMsg(code),
		}
	} //创建用户

	return &serializer.Response{
		Status: code,
		Msg:    pkg.GetMsg(code),
	} //成功
}

func (service *UserService) Login() serializer.Response {
	var user model.User
	code := pkg.SUCCESS
	if err := model.DB.Where("username=?", service.UserName).First(&user).Error; err != nil {
		// 如果查询不到，返回相应的错误
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			util.LogrusObj.Info(err)
			code = pkg.ErrorNotExistUser
			return serializer.Response{
				Status: code,
				Msg:    pkg.GetMsg(code),
			}
		}
		util.LogrusObj.Info(err)
		code = pkg.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    pkg.GetMsg(code),
		}
	}
	if !user.CheckPassword(service.Password) {
		code = pkg.ErrorNotCompare
		return serializer.Response{
			Status: code,
			Msg:    pkg.GetMsg(code),
		}
	}
	token, err := util.GenerateToken(user.ID, service.UserName, 0)
	if err != nil {
		util.LogrusObj.Info(err)
		code = pkg.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    pkg.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Data: serializer.TokenData{
			User:  serializer.BuildUser(user),
			Token: token,
		},
		Msg: pkg.GetMsg(code),
	}
}
