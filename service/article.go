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
	//ID      uint   `form:"id" json:"id"`
	Title   string `form:"title" json:"title" binding:"max=100"`
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
	err := model.DB.Model(&model.User{}).Where("uid = ?", id).First(&user).Error
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
	err = model.DB.Save(&article).Error
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

func (s *ShowArticle) Show() serializer.Response {
	code := pkg.SUCCESS
	var article model.Article
	var count int64
	err := model.DB.Model(&model.Article{}).First(&article).Count(&count).Error
	if err != nil {
		util.LogrusObj.Info(err)
		code = pkg.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    pkg.GetMsg(code),
			Error:  err.Error(),
		}
	}
	if count < 1 {
		return serializer.Response{
			Status: code,
			Msg:    pkg.GetMsg(code),
			Error:  "没有公告数据",
		}
	}
	var articles []model.Article
	model.DB.Model(&model.Article{}).Find(&articles)
	code = pkg.SUCCESS
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildArticles(articles),
		Msg:    pkg.GetMsg(code),
	}
}

func (u *UpdateArticle) Update(uid uint, id string) serializer.Response {
	code := pkg.SUCCESS
	var user model.User
	err := model.DB.Model(&model.User{}).Where("uid = ?", uid).First(&user).Error
	if err != nil {
		code = pkg.ErrorDatabase
		util.LogrusObj.Info(err)
		return serializer.Response{
			Status: code,
			Msg:    pkg.GetMsg(code),
			Error:  err.Error(),
		}
	}
	if user.Capacity != true {
		code = pkg.ErrorDatabase
		util.LogrusObj.Info(err)
		return serializer.Response{
			Status: code,
			Msg:    pkg.GetMsg(code),
			Error:  err.Error(),
		}
	}
	var article model.Article
	err = model.DB.Model(&model.Article{}).Where("id = ?", id).First(&article).Error
	if err != nil {
		code = pkg.ErrorDatabase
		util.LogrusObj.Info(err)
		return serializer.Response{
			Status: code,
			Msg:    pkg.GetMsg(code),
			Error:  err.Error(),
		}
	}
	if u.Title == "" {
	} else {
		article.Tile = u.Title
	}

	if u.Content != "" {
		article.Content = u.Content
	} else {
	}
	err = model.DB.Save(&article).Error
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
		Data:   serializer.BuildArticle(article),
		Msg:    pkg.GetMsg(code),
	}
}

func (d *DeleteArticle) DeleteArticle(uid uint, id string) serializer.Response {
	code := pkg.SUCCESS
	var user model.User
	err := model.DB.Model(&model.User{}).Where("uid = ?", uid).First(&user).Error
	if err != nil {
		code = pkg.ErrorDatabase
		util.LogrusObj.Info(err)
		return serializer.Response{
			Status: code,
			Msg:    pkg.GetMsg(code),
			Error:  err.Error(),
		}
	}
	if user.Capacity != true {
		code = pkg.ErrorDatabase
		util.LogrusObj.Info(err)
		return serializer.Response{
			Status: code,
			Msg:    pkg.GetMsg(code),
			Error:  err.Error(),
		}
	}
	var article model.Article
	if err := model.DB.Where("id = ?", id).First(&article).Error; err != nil {
		code = pkg.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    pkg.GetMsg(code),
			Error:  err.Error(),
		}
	}
	if err := model.DB.Delete(&article).Error; err != nil {
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
		Msg:    pkg.GetMsg(code),
	}
}
