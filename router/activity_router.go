package router

import (
	"github.com/danangkonang/todolist-app/config"
	"github.com/danangkonang/todolist-app/controller"
	"github.com/danangkonang/todolist-app/service"
	"github.com/gorilla/mux"
)

func ActivityRouter(router *mux.Router, db *config.DB) {
	c := controller.NewActivityController(
		service.NewServiceActivity(db),
	)
	router.HandleFunc("/activity-groups", c.FindActivities).Methods("GET")
	router.HandleFunc("/activity-groups/{id:[0-9]+}", c.FindActivitiyById).Methods("GET")
	router.HandleFunc("/activity-groups", c.CreateActivity).Methods("POST")
	router.HandleFunc("/activity-groups/{id:[0-9]+}", c.UpdateActivity).Methods("PATCH")
	router.HandleFunc("/activity-groups/{id:[0-9]+}", c.DeleteActivity).Methods("DELETE")
}
