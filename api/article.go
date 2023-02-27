package api

import (
	"github.com/gin-gonic/gin"
	"school_manager/service"
	"school_manager/util"
)

func CreateArticle(c *gin.Context) {
	createArticle := service.CreateArticle{}
	chaim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createArticle); err == nil {
		res := createArticle.Create(chaim.Uid)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ShowArticle(c *gin.Context) {
	showArticle := service.ShowArticle{}
	//chaim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&showArticle); err == nil {
		res := showArticle.Show()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func UpdateArticle(c *gin.Context) {
	updateArticle := service.UpdateArticle{}
	chaim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&updateArticle); err == nil {
		res := updateArticle.Update(chaim.Uid)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}
