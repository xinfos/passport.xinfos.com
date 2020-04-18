package model

import (
	"passport.xinfos.com/driver"

	"github.com/jinzhu/gorm"
)

//User - User Model struct
type User struct {
	ID       uint64 `json:"user_id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	IDCard   string `json:"id_card"`
	Birthday string `json:"birthday"`
	Gender   int    `json:"gender"`
	Age      int    `json:"age"`
}

var user *User

func UserModel() *User {
	return user
}

func (u *User) TableName() string {
	return `t_user`
}

//Create - create user
func (u *User) Create() error {
	if err := driver.DB.Table(u.TableName()).Create(&u).Error; err != nil {
		return err
	}
	return nil
}

//FindByID - 根据`user_id`查询用户信息
func (u *User) FindByID(userID uint64) (*User, error) {
	return u.findByMap(map[string]interface{}{
		"id": userID,
	})
}

//FindByIDCard - 根据`id_card`查询用户信息
func (u *User) FindByIDCard(idCard string) (*User, error) {
	return u.findByMap(map[string]interface{}{
		"id_card": idCard,
	})
}

//FindByPhone - 根据`phone`查询用户信息
func (u *User) FindByPhone(phone string) (*User, error) {
	return u.findByMap(map[string]interface{}{
		"phone": phone,
	})
}

func (u *User) findByMap(wheremaps map[string]interface{}) (*User, error) {
	var user User
	if err := driver.DB.Table(u.TableName()).Where(wheremaps).Find(&user).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}
	return &user, nil
}
