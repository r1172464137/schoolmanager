package service

import (
	"school_manager/model"
	"school_manager/pkg"
	"school_manager/serializer"
	"school_manager/util"
)

type ShowLeave struct {
}

type DeleteLeave struct {
}

type CreateLeave struct {
	Uid    uint   `json:"uid" form:"uid" binding:""`
	Name   string `json:"name" form:"name" binding:""`
	Reason string `json:"reason" form:"reason" binding:""`
	Time   uint   `json:"time" form:"time" binding:""`
}

type UpdateLeave struct {
	ID     uint `json:"id" form:"id" binding:""`
	Uid    uint `json:"uid" form:"uid" binding:""`
	Status uint `json:"status" form:"status" binding:""`
}

// 创建请假申请
func (createLeave *CreateLeave) Create(uid uint) serializer.Response {
	//var user model.User
	//model.DB.Model(&model.User{}).Where("uid = ?", id).First(&user)
	//util.LogrusObj.Info(uid)
	leave := model.Leave{
		Uid:    uid,
		Name:   createLeave.Name,
		Reason: createLeave.Reason,
		Time:   createLeave.Time,
		Status: 3,
	}
	code := pkg.SUCCESS
	err := model.DB.Create(&leave).Error
	if err != nil {
		util.LogrusObj.Info(err)
		code = pkg.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    pkg.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildLeave(leave),
		Msg:    pkg.GetMsg(code),
	}
}

func (s *ShowLeave) Show(uid uint) serializer.Response {
	code := pkg.SUCCESS
	var user model.User
	model.DB.Model(&model.User{}).Where("uid = ?", uid).First(&user)
	if user.Capacity != true { //老师，可以查看所有自己学院下的信息
		var studentLeave []model.Leave
		err := model.DB.Model(&model.Leave{}).Where("uid = ?", uid).Find(&studentLeave).Error
		if err != nil {
			code = pkg.ErrorDatabase
			util.LogrusObj.Info(err)
			return serializer.Response{
				Status: code,
				Msg:    pkg.GetMsg(code),
				Error:  err.Error(),
			}
		}
		code = pkg.SUCCESS
		return serializer.Response{
			Status: code,
			Data:   serializer.BuildLeaves(studentLeave),
			Msg:    pkg.GetMsg(code),
		}
	}
	var leaves []model.Leave
	err := model.DB.Model(&model.Leave{}).Find(&leaves).Error
	if err != nil {
		util.LogrusObj.Info(err)
		code = pkg.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    pkg.GetMsg(code),
			Error:  err.Error(),
		}
	}
	code = pkg.SUCCESS
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildLeaves(leaves),
		Msg:    pkg.GetMsg(code),
	}

}

// 更新申请（其实就是批准申请）
func (u *UpdateLeave) Update(uid uint) serializer.Response {
	code := pkg.SUCCESS
	var user model.User
	model.DB.Model(&model.User{}).Where("uid = ?", uid).First(&user)
	if user.Capacity == false { //身份为学生
		code = pkg.ErrorDatabase
		return serializer.Response{
			Status: code,
			//Data:   nil,
			Msg:   pkg.GetMsg(code),
			Error: "error user in use",
		}
	}
	//var leave model.Leave
	//leave.Status = u.Status
	err := model.DB.Model(&model.Leave{}).Where("id = ?", u.ID).Update("status", u.Status).Error
	if err != nil {
		code = pkg.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    pkg.GetMsg(code),
			Error:  err.Error(),
		}
	}
	//leave.Status = u.Status
	//model.DB.Save(&leave)
	code = pkg.SUCCESS
	return serializer.Response{
		Status: code,
		Data:   nil,
		Msg:    "更新成功",
	}
}
