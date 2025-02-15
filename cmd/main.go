package main

import (
	"github.com/dmitrycr/todo-app"
	"github.com/dmitrycr/todo-app/pkg/handler"
	"github.com/dmitrycr/todo-app/pkg/repository"
	"github.com/dmitrycr/todo-app/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error run server: %s", err.Error())
	}
}
