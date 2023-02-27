package service

import (
	"school_manager/model"
	"school_manager/pkg"
	"school_manager/serializer"
	"school_manager/util"
)

// 删除公告的服务
type DeleteArticle struct {
}

// 删除公告的服务
type ShowArticle struct {
}

// 更新公告的服务
type UpdateArticle struct {
	Uid     uint   `form:"uid" json:"uid"`
	Title   string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Content string `form:"content" json:"content" binding:"max=1000"`
	//Status  int    `form:"status" json:"status"` // 0 待办   1已完成
}

// 创建公告的服务
type CreateArticle struct {
	Title   string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Content string `form:"content" json:"content" binding:"max=1000"`
	//Status  int    `form:"status" json:"status"` // 0 待办   1已完成
}

func (c *CreateArticle) Create(id uint) serializer.Response {
	code := pkg.SUCCESS
	var user model.User
	err := model.DB.Model(&model.User{}).Where("uid", id).First(&user).Error
	if err != nil {
		util.LogrusObj.Info(err)
		code = pkg.ErrorNotExistUser
		return serializer.Response{
			Status: code,
			Msg:    pkg.GetMsg(code),
			Error:  err.Error(),
		}
	}
	if user.Capacity != true {
		code = pkg.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    pkg.GetMsg(code),
			Error:  "学生不允许操作数据库",
		}
	}
	var article model.Article
	article.Tile = c.Title
	article.Content = c.Content
	article.Publisher = id
	err = model.DB.Save(&user).Error
	if err != nil {
		util.LogrusObj.Info(err)
		code = pkg.ErrorDatabase
		return serializer.Response{
			Status: code,
			//Data:   nil,
			Msg:   pkg.GetMsg(code),
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildArticle(article),
		Msg:    pkg.GetMsg(code),
	}
}
