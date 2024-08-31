package routerModule

import (
	"easy-go-admin/app/controller"
	"easy-go-admin/app/middleware"
	"github.com/gin-gonic/gin"
)

// SetupCommonRouter 设置全局路由
func SetupCommonRouter(base *gin.RouterGroup) {
	index := base.Group("/index")
	index.GET("/index", controller.Index)

	api := base.Group("/api")

	account := api.Group("/account")
	account.POST("/login", controller.Login)
	account.POST("/reg", controller.Reg)
	// 鉴权中间件
	account.Use(middleware.AuthMiddleware())
	{
		account.GET("/getDetail", controller.GetAccountDetail)
	}
}
