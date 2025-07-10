package main

import (
	"log"
	"net/http"

	gen "github.com/RicardoTlatelpa/uniqueidgen"
)

var (
	idGen *gen.Gen	
	baseURL = "http://localhost:8080/"
)

func main() {
	var err error
	idGen, err = gen.NewGen(1)
	if err != nil {
		log.Fatalf("failed to create generator: %v", err)
	}

	http.HandleFunc("/shorten", handleShorten)
	http.HandleFunc("/", handleRedirect)

	log.Println("server started at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
