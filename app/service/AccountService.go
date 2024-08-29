package service

import (
	"easy-go-admin/app/dao"
	"easy-go-admin/app/entity"
	"easy-go-admin/config/jwt"
	"easy-go-admin/config/message"
	"easy-go-admin/config/util"
	"github.com/gin-gonic/gin"
	Validate "github.com/go-playground/validator/v10"
)

// IAccountService 定义接口
type IAccountService interface {
	Login(c *gin.Context, dto account_entity.LoginData)
	Reg(c *gin.Context, dto account_entity.RegData)
	GetDetail(c *gin.Context)
	SendCaptcha(c *gin.Context, dto account_entity.SendCaptcha)
}

type accountServiceImpl struct{}

func AccountService() IAccountService {
	return &accountServiceImpl{}
}

// SendCaptcha 发送验证码
func (s *accountServiceImpl) SendCaptcha(c *gin.Context, dto account_entity.SendCaptcha) {
	//TODO implement me
	panic("implement me")
}

// Login 用户登陆
func (s *accountServiceImpl) Login(c *gin.Context, dto account_entity.LoginData) {
	// 登陆参数验证
	err := Validate.New().Struct(dto)
	if err != nil {
		message.Fail(c, int(message.Code.BADREQUEST), message.Code.GetMessage(message.Code.BADREQUEST))
		return
	}
	// 验证用户信息
	Account, err := account_dao.AccountDetailByUsername(dto.UserName)
	if err != nil || Account.ID == 0 {
		message.Fail(c, int(message.Code.NOTFOUND), "用户名不存在")
		return
	}
	if Account.Password != util.EncryptionPassword(dto.Password, Account.Salt) {
		message.Fail(c, int(message.Code.PASSWORDNOTTRUE), message.Code.GetMessage(message.Code.PASSWORDNOTTRUE))
		return
	}
	const status int = 0
	if Account.Status == status {
		message.Fail(c, int(message.Code.STATUSISENABLE), message.Code.GetMessage(message.Code.STATUSISENABLE))
		return
	}

	// 生成token
	tokenString, _ := jwt.GenerateTokenByAdmin(Account)
	message.Success(c, map[string]interface{}{"token": tokenString, "Account": Account})
}

// Reg 注册用户
func (s *accountServiceImpl) Reg(c *gin.Context, dto account_entity.RegData) {
	// 创建参数验证
	err := Validate.New().Struct(dto)
	if err != nil {
		message.Fail(c, int(message.Code.BADREQUEST), message.Code.GetMessage(message.Code.BADREQUEST))
		return
	}
	Account, err := account_dao.AccountDetailByUsername(dto.UserName)
	if err == nil && Account.ID != 0 {
		message.Fail(c, int(message.Code.BADREQUEST), "用户名已被占用")
		return
	}
	salt := util.GenerateRandomString(15)
	pass := util.EncryptionPassword(dto.Password, salt)
	var avatar string
	if dto.Avatar != nil {
		avatar = *dto.Avatar
	} else {
		avatar = ""
	}
	if dto.Code != 6666 {
		message.Fail(c, int(message.Code.BADREQUEST), "验证码不正确")
		return
	}
	createData := account_entity.Account{
		UserName: dto.UserName,
		Password: pass,
		Salt:     salt,
		Email:    dto.Email,
		Avatar:   avatar,
		Mobile:   dto.Mobile,
	}
	err = account_dao.AccountCreate(createData)
	if err != nil {
		message.Fail(c, int(message.Code.CREATERESOURCEFAILED), "创建用户失败")
		return
	}
	message.Success(c, nil)
}

// GetDetail 获取当前用户信息
func (s *accountServiceImpl) GetDetail(c *gin.Context) {
	user, err := jwt.GetAccountInfo(c)
	if err != nil {
		message.Fail(c, int(message.Code.FORBIDDEN), err.Error())
		return
	}

	response := map[string]any{
		"id":       user.Id,
		"username": user.UserName,
		"avatar":   user.Avatar,
		"email":    user.Email,
		"mobile":   user.Mobile,
	}
	message.Success(c, response)
}
