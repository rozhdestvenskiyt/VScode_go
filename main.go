package main

import (
	"fmt"
	"itprogergolang/internal/user"
	"itprogergolang/pkg/logging"
	"log"
	"net"
	"time"

	"net/http"

	"github.com/julienschmidt/httprouter"
)


func main() {
	logger := logging.GetLogger()
	logger.Info("Проекта начался")
	router := httprouter.New()
	handler := user.Handler_user{}
	handler.Register(router)

	start_server(router , &logger)
}

func start_server(router *httprouter.Router , logger *logging.Logger) {
	listener , err := net.Listen("tcp" , ":8080")
	if err != nil {
		log.Fatal("Error with listener")
	}
	logger.Info("Запустили listener")


	server := &http.Server{
		Handler: router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}
	
	log.Fatal(server.Serve(listener))
	fmt.Println("Project died")
}