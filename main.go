package main

import (
	"github.com/danangkonang/todolist-app/app"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	app.Run()
}
