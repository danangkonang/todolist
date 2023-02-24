package router

import (
	"github.com/danangkonang/todolist-app/config"
	"github.com/danangkonang/todolist-app/controller"
	"github.com/danangkonang/todolist-app/service"
	"github.com/gorilla/mux"
)

func TodoRouter(router *mux.Router, db *config.DB) {
	c := controller.NewTodoController(
		service.NewServiceTodo(db),
		service.NewServiceActivity(db),
	)
	router.HandleFunc("/todo-items", c.FindTodos).Methods("GET")
	router.HandleFunc("/todo-items/{id:[0-9]+}", c.FindTodoById).Methods("GET")
	router.HandleFunc("/todo-items", c.CreateTodo).Methods("POST")
	router.HandleFunc("/todo-items/{id:[0-9]+}", c.UpdateTodo).Methods("PATCH")
	router.HandleFunc("/todo-items/{id:[0-9]+}", c.DeleteTodo).Methods("DELETE")
}
