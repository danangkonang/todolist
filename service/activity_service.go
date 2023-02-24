package service

import (
	"github.com/danangkonang/todolist-app/config"
	"github.com/danangkonang/todolist-app/model"
)

type ServiceActivity interface {
	FindActivities() ([]*model.Activity, error)
	CreateActivity(c *model.Activity) (*model.Activity, error)
	FindActivitiyById(id int) (*model.Activity, error)
	UpdateActivity(c *model.Activity) (*model.Activity, error)
	DeleteActivity(id int) error
}

func NewServiceActivity(Con *config.DB) ServiceActivity {
	return &connection{
		Mysql: Con.Mysql,
	}
}

func (r *connection) FindActivities() ([]*model.Activity, error) {
	fac := make([]*model.Activity, 0)
	err := r.Mysql.Find(&fac).Error
	if err != nil {
		return nil, err
	}
	return fac, nil
}

func (r *connection) FindActivitiyById(id int) (*model.Activity, error) {
	fac := new(model.Activity)
	err := r.Mysql.First(&fac, "activity_id=?", id).Error
	if err != nil {
		return nil, err
	}
	return fac, nil
}

func (r *connection) CreateActivity(c *model.Activity) (*model.Activity, error) {
	err := r.Mysql.Save(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *connection) UpdateActivity(c *model.Activity) (*model.Activity, error) {
	fac := new(model.Activity)
	err := r.Mysql.First(fac, "activity_id=?", c.Id).Error
	if err != nil {
		return nil, err
	}
	err = r.Mysql.Model(fac).Update("title", c.Title).Error
	if err != nil {
		return nil, err
	}
	return fac, nil
}

func (r *connection) DeleteActivity(id int) error {
	fac := new(model.Activity)
	err := r.Mysql.First(fac, "activity_id=?", id).Error
	if err != nil {
		return err
	}
	err = r.Mysql.Where("activity_id=?", id).Delete(fac).Error
	if err != nil {
		return err
	}
	return nil
}
