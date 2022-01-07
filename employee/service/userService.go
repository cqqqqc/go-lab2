package service

import (
	"Lab2/common/entity"
	"Lab2/user/dao"
)

func CreateUser(user *entity.User) (err error) {
	if err = dao.Db.Create(user).Error; err != nil {
		return err
	}
	return
}
