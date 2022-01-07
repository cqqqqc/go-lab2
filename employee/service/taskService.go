package service

import (
	"Lab2/common/entity"
	"Lab2/task/dao"
)

func CreateTask(task *entity.Task) (err error) {
	if err = dao.Db.Create(task).Error; err != nil {
		return err
	}
	return
}
