package main

// go run main.go handlers.go routes.go repo.go user.go warning.go

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))

}
