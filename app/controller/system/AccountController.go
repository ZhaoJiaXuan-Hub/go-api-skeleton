package controllerSystem

import (
	"easy-go-admin/app/entity"
	"easy-go-admin/app/service"
	"github.com/gin-gonic/gin"
)

// Login 用户登陆
// @Summary 用户登陆接口
// @Produce json
// @Description 前台用户登陆接口
// @Param data body account_entity.LoginData true "data"
// @Success 200 {object} message.Response
// @router /api/system/account/login [post]
func Login(c *gin.Context) {
	var dto account_entity.LoginData
	_ = c.BindJSON(&dto)
	service.AccountService().Login(c, dto)
}

// Reg 注册用户
// @Summary 用户注册接口
// @Produce json
// @Description 前台用户注册接口
// @Param data body account_entity.RegData true "data"
// @Success 200 {object} message.Response
// @router /api/system/account/reg [post]
func Reg(c *gin.Context) {
	var dto account_entity.RegData
	_ = c.BindJSON(&dto)
	service.AccountService().Reg(c, dto)
}

// GetAccountDetail 获取登录用户信息
// @Summary 获取登录用户信息
// @Produce json
// @Description 后台获取登录用户信息
// @Security ApiKeyAuth
// @Success 200 {object} message.Response
// @router /api/system/account/getDetail [get]
func GetAccountDetail(c *gin.Context) {
	service.AccountService().GetDetail(c)
}

func SendCaptcha(c *gin.Context) {
	var dto account_entity.SendCaptcha
	_ = c.BindJSON(&dto)
	service.AccountService().SendCaptcha(c, dto)
}
