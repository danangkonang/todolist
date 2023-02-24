package service

import (
	"github.com/danangkonang/todolist-app/config"
	"github.com/danangkonang/todolist-app/model"
)

type ServiceTodo interface {
	FindTodos(activity_group_id int) ([]*model.Todo, error)
	CreateTodo(c *model.Todo) (*model.Todo, error)
	FindTodoById(id int) (*model.Todo, error)
	UpdateTodo(c *model.Todo) (*model.Todo, error)
	DeleteTodo(t *model.Todo) error
}

func NewServiceTodo(Con *config.DB) ServiceTodo {
	return &connection{
		Mysql: Con.Mysql,
	}
}

func (r *connection) FindTodos(activity_group_id int) ([]*model.Todo, error) {
	fac := make([]*model.Todo, 0)
	base := r.Mysql
	if activity_group_id > 0 {
		base = r.Mysql.Where("activity_group_id=?", activity_group_id)
	}
	err := base.Find(&fac).Error
	if err != nil {
		return nil, err
	}
	return fac, nil
}

func (r *connection) FindTodoById(id int) (*model.Todo, error) {
	fac := new(model.Todo)
	err := r.Mysql.First(&fac, "todo_id=?", id).Error
	if err != nil {
		return nil, err
	}
	return fac, nil
}

func (r *connection) CreateTodo(c *model.Todo) (*model.Todo, error) {
	err := r.Mysql.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *connection) UpdateTodo(c *model.Todo) (*model.Todo, error) {
	fac := new(model.Todo)
	err := r.Mysql.First(fac, "todo_id=?", c.Id).Error
	if err != nil {
		return nil, err
	}

	err = r.Mysql.Model(fac).Updates(map[string]interface{}{"title": c.Title, "priority": c.Priority, "is_active": c.IsActive}).Error
	if err != nil {
		return nil, err
	}
	return fac, nil
}

func (r *connection) DeleteTodo(t *model.Todo) error {
	fac := new(model.Todo)
	err := r.Mysql.First(fac, "todo_id=? AND title=?", t.Id, t.Title).Error
	if err != nil {
		return err
	}
	err = r.Mysql.Where("todo_id=?", t.Id).Delete(fac).Error
	if err != nil {
		return err
	}
	return nil
}
