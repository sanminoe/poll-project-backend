package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	pollRoutes "github.com/sanminoe/poll-project-backend/pkg/routes"
)

func main() {

	router := mux.NewRouter()

	pollRoutes.RegisterPollRoutes(router)

	http.Handle("/", router)
	fmt.Println("Server started")
	log.Fatal(http.ListenAndServe("localhost:9010", router))

}
