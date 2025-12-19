package main

import (

	"itprogergolang/internal/user"
	"log"
	"net"
	"time"

	"net/http"

	"github.com/julienschmidt/httprouter"
)


func main() {
	router := httprouter.New()
	handler := user.Handler_user{}
	handler.Register(router)

	start_server(router)
}

func start_server(router *httprouter.Router) {
	listener , err := net.Listen("tcp" , ":8080")
	if err != nil {
		log.Fatal("Error with listener")
	}

	server := &http.Server{
		Handler: router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	log.Fatal(server.Serve(listener))
}