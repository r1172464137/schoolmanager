package api

import (
	"school_manager/service"
	"school_manager/util"

	"github.com/gin-gonic/gin"
)

func CreateLeave(c *gin.Context) {
	createLeave := service.CreateLeave{}
	chaim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createLeave); err == nil {
		res := createLeave.Create(uint(chaim.Uid))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ShowLeave(c *gin.Context) {
	showLeave := service.ShowLeave{}
	chaim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&showLeave); err == nil {
		res := showLeave.Show(chaim.Uid)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func UpdateLeave(c *gin.Context) {
	updateLeave := service.UpdateLeave{}
	id := c.Param("id")
	chaim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&updateLeave); err == nil {
		res := updateLeave.Update(chaim.Uid, id)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}
