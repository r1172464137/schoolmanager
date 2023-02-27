package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"school_manager/api"
	"school_manager/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Use(sessions.Sessions("mysession", store))
	r.Use(middleware.Cors())
	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})
		// 用户操作
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)
		//todo:简化注册逻辑
		auth := v1.Group("/")
		auth.Use(middleware.JWT())
		{
			auth.POST("leave", api.CreateLeave)
			auth.GET("leave", api.ShowLeave)
			auth.PUT("leave", api.UpdateLeave)
		}
	}
	return r
}
