package dao

import (
	"easy-go-admin/app/constant"
	"easy-go-admin/app/entity"
	"easy-go-admin/config"
	"easy-go-admin/config/database"
)

// AccountDetailByUsername 通过用户名获取用户详细
func AccountDetailByUsername(username string) (systemAccount entity.Account, err error) {
	if err := database.Db.Where("user_name = ?", username).First(&systemAccount).Error; err != nil {
		return systemAccount, err // 或者根据具体情况返回错误或默认值
	}
	return systemAccount, nil
}

// AccountDetailById 通过ID获取用户详细
func AccountDetailById(id uint) (account entity.Account, err error) {
	if err := database.Db.Preload("Roles.Menus.RolesInfo").Preload("Dept").Where("id = ?", id).First(&account).Error; err != nil {
		return account, err // 或者根据具体情况返回错误或默认值
	}
	// 检查并设置默认头像
	if account.Avatar == "" {
		account.Avatar = config.Config.FileSettings.Host + constant.DefaultAvatarUrl // 设置默认头像 URL
	}
	return account, nil
}

// AccountCreate 新增用户
func AccountCreate(account entity.Account) (err error) {
	err = database.Db.Create(&account).Error
	return err
}
