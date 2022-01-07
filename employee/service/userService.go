package service

import (
	"Lab2/common/entity"
	"Lab2/user/dao"
)

func CreateUser(user *entity.User) (err error) {
	a := dao.Db
	if err = a.Create(user).Error; err != nil {
		return err
	}
	return
}
