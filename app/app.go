package app

import (
	"fmt"
	"net/http"
	"os"

	"github.com/danangkonang/todolist-app/config"
	"github.com/danangkonang/todolist-app/helper"
	"github.com/danangkonang/todolist-app/model"
	"github.com/danangkonang/todolist-app/router"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Run() {
	r := mux.NewRouter().StrictSlash(false)
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		helper.MakeRespon(w, 404, "page not found", nil)
	})

	r.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		helper.MakeRespon(w, http.StatusMethodNotAllowed, "Method NotAllowed", nil)
	})
	db := config.Connection(
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DBNAME"),
	)
	db.Mysql.AutoMigrate(
		&model.Activity{},
		&model.Todo{},
	)
	router.ActivityRouter(r, db)
	router.TodoRouter(r, db)
	serverloging := fmt.Sprintf("local server started at http://localhost:%s", os.Getenv("APP_PORT"))
	fmt.Println(serverloging)

	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("APP_PORT")), handlers.CORS(
		handlers.AllowedHeaders([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PATCH", "DELETE"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowCredentials(),
	)(r))
}
