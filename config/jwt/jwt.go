package jwt

import (
	"easy-go-admin/app/constant"
	"easy-go-admin/app/entity"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

type userStdClaims struct {
	account_entity.JwtAccountData
	jwt.StandardClaims
}

// TokenExpireDuration 过期时间
const TokenExpireDuration = time.Hour * 2

// Secret 密钥
const Secret = "easyadmin"

var (
	ErrAbsent      = "令牌不存在"
	ErrInvalid     = "令牌无效"
	ErrExpired     = "令牌已过期"
	ErrNotValidYet = "令牌尚未生效"
)

// GenerateTokenByAdmin 生成用户token
func GenerateTokenByAdmin(account account_entity.Account) (string, error) {
	var jwtAdmin = account_entity.JwtAccountData{
		Id:       account.ID,
		UserName: account.UserName,
		Mobile:   account.Mobile,
		Avatar:   account.Avatar,
		Email:    account.Email,
	}
	c := userStdClaims{
		jwtAdmin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "admin",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(Secret))
}

// ValidateToken 解析token
func ValidateToken(tokenString string) (*account_entity.JwtAccountData, error) {
	if len(tokenString) == 0 {
		return nil, errors.New(ErrAbsent)
	}

	claims := userStdClaims{}
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(Secret), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			switch {
			case ve.Errors == jwt.ValidationErrorExpired:
				return nil, errors.New(ErrExpired)
			case ve.Errors == jwt.ValidationErrorNotValidYet:
				return nil, errors.New(ErrNotValidYet)
			default:
				return nil, errors.New(ErrInvalid)
			}
		}
		return nil, err
	}

	return &claims.JwtAccountData, nil
}

// GetAccountID 返回用户id
func GetAccountID(c *gin.Context) (uint, error) {
	u, exist := c.Get(constant.ContextKeyUserObj)
	if !exist {
		return 0, errors.New("未找到id")
	}
	admin, ok := u.(account_entity.Account)
	if !ok {
		return admin.ID, nil
	}
	return 0, errors.New("未找到id")
}

// GetAccountInfo 返回用户信息
func GetAccountInfo(c *gin.Context) (*account_entity.JwtAccountData, error) {
	u, exist := c.Get(constant.ContextKeyUserObj)
	if !exist {
		return nil, errors.New("未找到用户")
	}

	// 确保从上下文中获取的数据类型是 *account_entity.JwtAccountData
	admin, ok := u.(*account_entity.JwtAccountData)
	if !ok {
		return nil, errors.New("类型错误：未找到正确的用户对象")
	}

	return admin, nil
}
