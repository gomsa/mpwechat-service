package service

import (
	mp "github.com/gomsa/mpwechat-service/proto/wechat"

	"github.com/jinzhu/gorm"
)

//Repository 仓库接口
type Repository interface {
	Create(user *mp.User) error
	GetByOpenid(openid string) (*mp.User, error)
}

// UserRepository 用户仓库
type UserRepository struct {
	DB *gorm.DB
}

// Create 创建用户
func (repo *UserRepository) Create(user *mp.User) error {
	if err := repo.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// GetByOpenid 根据 openid 获取用户信息
func (repo *UserRepository) GetByOpenid(openid string) (*mp.User, error) {
	user := &mp.User{}
	if err := repo.DB.Where("openid = ?", openid).
		First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
