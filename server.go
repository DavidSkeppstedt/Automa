package main

import (
	"github.com/DavidSkeppstedt/Automa/api"
	"log"
	"net/http"
)

func main() {
	r := api.NewRouter()
	log.Println("Server started")
	http.ListenAndServe(":8080", r)
}
