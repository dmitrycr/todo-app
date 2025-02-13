package main

import (
	"github.com/dmitrycr/todo-app"
	"log"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error run server: %s", err.Error())
	}
}
