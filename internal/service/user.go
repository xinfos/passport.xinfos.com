package service

import (
	"passport.xinfos.com/internal/model"
	"passport.xinfos.com/internal/repository"

	"github.com/pkg/errors"
)

const (
	module = "service.user"
)

//UserService - 用户服务
type UserService struct {
	repo *repository.UserRepository
}

//NewUserService - 实例化用户服务
func NewUserService() *UserService {
	return &UserService{
		repo: repository.NewUserRepository(),
	}
}

//GetUserInfoByID - get
func (s *UserService) GetUserInfoByID(userID uint64) (*model.User, error) {
	user, err := s.repo.FindByID(userID)
	if user == nil {
		if err == nil {
			err = errors.New("当前用户信息不存在")
		}
		return nil, errors.Wrap(err, "当前用户信息不存在")
	}
	return user, nil
}

func (s *UserService) Create(u *model.User) (uint64, error) {
	//1、当前`IDcard`已存在
	user, _ := s.repo.FindByIDCard(u.IDCard)
	if user != nil && user.ID > 0 {
		return 0, errors.Wrap(errors.New("当前用户信息已存在"), "当前用户信息已存在")
	}

	//2、当前`Phone`已存在
	user, _ = s.repo.FindByPhone(u.Phone)
	if user != nil && user.ID > 0 {
		return 0, errors.Wrap(errors.New("当前用户信息已存在"), "当前用户信息已存在")
	}

	userId, err := s.repo.Create(u)
	if err != nil {
		return 0, errors.Wrap(err, "创建失败")
	}
	return userId, nil
}
