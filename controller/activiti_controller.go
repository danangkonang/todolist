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

type activityController struct {
	service service.ServiceActivity
}

func NewActivityController(a service.ServiceActivity) *activityController {
	return &activityController{
		service: a,
	}
}

func (c *activityController) FindActivities(w http.ResponseWriter, r *http.Request) {
	res, err := c.service.FindActivities()
	if err != nil {
		helper.MakeRespon(w, 400, "Internal server error", nil)
		return
	}
	helper.MakeRespon(w, 200, "Success", res)
}

func (c *activityController) FindActivitiyById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	res, err := c.service.FindActivitiyById(id)
	switch {
	case err == gorm.ErrRecordNotFound:
		helper.MakeRespon(w, 404, fmt.Sprintf("Activity with ID %d Not Found", id), nil)
		return
	case err != nil:
		helper.MakeRespon(w, 500, "Internal server error", nil)
		return
	}
	helper.MakeRespon(w, 200, "Success", res)
}

func (c *activityController) CreateActivity(w http.ResponseWriter, r *http.Request) {
	var p model.Activity
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		helper.MakeRespon(w, 400, "Invalid json format", nil)
		return
	}
	res, err := c.service.CreateActivity(&p)
	if err != nil {
		helper.MakeRespon(w, 400, "Internal server error", nil)
		return
	}
	helper.MakeRespon(w, 200, "Success", res)
}

func (c *activityController) UpdateActivity(w http.ResponseWriter, r *http.Request) {
	var p model.Activity
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		helper.MakeRespon(w, 400, "Invalid json format", nil)
		return
	}
	p.Id = id
	res, err := c.service.UpdateActivity(&p)
	switch {
	case err == gorm.ErrRecordNotFound:
		helper.MakeRespon(w, 404, fmt.Sprintf("Activity with ID %d Not Found", id), nil)
		return
	case err != nil:
		helper.MakeRespon(w, 500, "Internal server error", nil)
		return
	}
	helper.MakeRespon(w, 200, "Success", res)
}

func (c *activityController) DeleteActivity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	err := c.service.DeleteActivity(id)
	switch {
	case err == gorm.ErrRecordNotFound:
		helper.MakeRespon(w, 404, fmt.Sprintf("Activity with ID %d Not Found", id), nil)
		return
	case err != nil:
		helper.MakeRespon(w, 500, "Internal server error", nil)
		return
	}
	helper.MakeRespon(w, 200, "Success", nil)
}
