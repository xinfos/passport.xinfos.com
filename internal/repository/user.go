package repository

import (
	"errors"
	"passport.xinfos.com/internal/model"
	"passport.xinfos.com/internal/repository/cache"
)

//UserRepository - 用户仓库
type UserRepository struct {
	c *cache.UserCache
}

//NewUserRepository - 初始化 UserRepository
func NewUserRepository() *UserRepository {
	return &UserRepository{
		c: cache.NewUserCache(),
	}
}
func (repo *UserRepository) Create(user *model.User) (uint64, error) {
	err := user.Create()
	if err != nil {
		return 0, err
	}
	if user == nil || user.ID <= 0 {
		return 0, errors.New("create fail")
	}
	repo.c.Set(user)

	return user.ID, nil
}

//FindByID - 根据ID获取用户信息
func (repo *UserRepository) FindByID(userID uint64) (*model.User, error) {
	data := repo.c.Get(userID)
	if data != nil && data.ID > 0 {
		return data, nil
	}

	data, _ = model.UserModel().FindByID(userID)

	if data != nil && data.ID == userID {
		repo.c.Set(data)
	}
	return data, nil
}

func (repo *UserRepository) FindByIDCard(idCard string) (*model.User, error) {
	data, err := model.UserModel().FindByIDCard(idCard)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//FindByPhone - 查询用户信息
func (repo *UserRepository) FindByPhone(phone string) (*model.User, error) {
	data, err := model.UserModel().FindByPhone(phone)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *UserRepository) FindAll() (*model.User, error) {

	return nil, nil
}
