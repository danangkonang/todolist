package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/danangkonang/todolist-app/helper"
	"github.com/danangkonang/todolist-app/model"
	"github.com/danangkonang/todolist-app/service"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type todoController struct {
	todo     service.ServiceTodo
	activity service.ServiceActivity
}

func NewTodoController(t service.ServiceTodo, a service.ServiceActivity) *todoController {
	return &todoController{
		todo:     t,
		activity: a,
	}
}

func (c *todoController) FindTodos(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	ids := v.Get("activity_group_id")
	activity_group_id, _ := strconv.Atoi(ids)
	res, err := c.todo.FindTodos(activity_group_id)
	if err != nil {
		helper.MakeRespon(w, 400, "Internal server error", nil)
		return
	}
	helper.MakeRespon(w, 200, "Success", res)
}

func (c *todoController) FindTodoById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	res, err := c.todo.FindTodoById(id)
	switch {
	case err == gorm.ErrRecordNotFound:
		helper.MakeRespon(w, 404, fmt.Sprintf("Todo with ID %d Not Found", id), nil)
		return
	case err != nil:
		helper.MakeRespon(w, 500, "Internal server error", nil)
		return
	}
	helper.MakeRespon(w, 200, "Success", res)
}

func (c *todoController) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var p model.Todo
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		helper.MakeRespon(w, 400, "Invalid json format", nil)
		return
	}
	_, err := c.activity.FindActivitiyById(p.ActivityGroupId)
	switch {
	case err == gorm.ErrRecordNotFound:
		helper.MakeRespon(w, 404, fmt.Sprintf("Activity with ID %d Not Found", p.ActivityGroupId), nil)
		return
	case err != nil:
		helper.MakeRespon(w, 500, "Internal server error", nil)
		return
	}

	res, err := c.todo.CreateTodo(&p)
	if err != nil {
		helper.MakeRespon(w, 400, "Internal server error", nil)
		return
	}
	helper.MakeRespon(w, 200, "Success", res)
}

func (c *todoController) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	var p model.Todo
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		fmt.Println(err.Error())
		helper.MakeRespon(w, 400, "Invalid json format", nil)
		return
	}
	p.Id = id
	res, err := c.todo.UpdateTodo(&p)
	switch {
	case err == gorm.ErrRecordNotFound:
		helper.MakeRespon(w, 404, fmt.Sprintf("Todo with ID %d Not Found", id), nil)
		return
	case err != nil:
		helper.MakeRespon(w, 500, "Internal server error", nil)
		return
	}
	helper.MakeRespon(w, 200, "Success", res)
}

func (c *todoController) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	var p model.Todo
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		fmt.Println(err.Error())
		helper.MakeRespon(w, 400, "Invalid json format", nil)
		return
	}
	p.Id = id
	err := c.todo.DeleteTodo(&p)
	switch {
	case err == gorm.ErrRecordNotFound:
		helper.MakeRespon(w, 404, fmt.Sprintf("Todo with ID:%d and Title:%s Not Found", id, p.Title), nil)
		return
	case err != nil:
		helper.MakeRespon(w, 500, "Internal server error", nil)
		return
	}
	helper.MakeRespon(w, 200, "Success", nil)
}
