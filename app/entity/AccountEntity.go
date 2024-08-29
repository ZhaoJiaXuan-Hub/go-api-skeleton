package account_entity

import (
	"easy-go-admin/config/util"
	"gorm.io/gorm"
)

func (Account) TableName() string {
	return "account"
}

// Account 用户
type Account struct {
	ID        uint           `gorm:"primaryKey;autoIncrement;comment:主键" json:"id"`
	UserName  string         `gorm:"not null;comment:用户名" json:"username"`
	Avatar    string         `gorm:"comment:头像;default:NULL" json:"avatar"`
	Password  string         `gorm:"not null;comment:密码" json:"password"`
	Salt      string         `gorm:"not null;comment:密码盐" json:"salt"`
	Email     string         `gorm:"comment:邮箱" json:"email"`
	Mobile    int            `gorm:"comment:手机号" json:"mobile"`
	Status    int            `gorm:"comment:状态;default:1" json:"status"`
	Balance   float64        `gorm:"comment:余额;default:0.00" json:"balance"`
	DeletedAt gorm.DeletedAt `gorm:"index;softDelete;comment:删除时间;default:NULL" json:"deleted_at"` // 软删除
	CreatedAt util.Htime     `gorm:"comment:创建时间;type:datetime" json:"created_at"`
	UpdatedAt util.Htime     `gorm:"comment:更新时间;type:datetime" json:"updated_at"`
}

type JwtAccountData struct {
	Id       uint   `json:"id"`
	UserName string `json:"username"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
	Mobile   int    `json:"mobile"`
}

type LoginData struct {
	UserName string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegData struct {
	UserName string  `json:"username" validate:"required"`
	Avatar   *string `json:"avatar,omitempty" validate:"omitempty"`
	Email    string  `json:"email" validate:"required"`
	Mobile   int     `json:"mobile" validate:"required"`
	Code     int     `json:"code" validate:"required"`
	Password string  `json:"password" validate:"required"`
}

type SendCaptcha struct {
	Mobile int `json:"mobile" validate:"required"`
}
