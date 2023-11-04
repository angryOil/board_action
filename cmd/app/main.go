package main

import (
	"board_action/cmd/app/handler"
	"board_action/internal/controller"
	"board_action/internal/repository"
	"board_action/internal/repository/infla"
	"board_action/internal/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	h := getHandler()

	r.PathPrefix("/board-actions").Handler(h)
	err := http.ListenAndServe(":8088", r)
	if err != nil {
		log.Panic("server listen err: ", err)
	}
}

func getHandler() http.Handler {
	return handler.NewHandler(controller.NewController(service.NewService(repository.NewRepository(infla.NewDB()))))
}
