package main

import (
	"log"
	"net/http"
	"sync"
	"github.com/RicardoTlatelpa/uniqueidgen"
)

var (
	idGen *gen.Gen
	store = make(map[string]string)
	mu = sync.RWMutex{}
	baseURL = "http://localhost:8080/"
)

func main() {
	var err error
	idGen, err = gen.NewGen(1)
	if err != nil {
		log.Fatalf("failed to create generator: %v", err)
	}

	//http.HandleFunc("/shorten", handleShorten)
	//http.HandleFunc("/", handleRedirect)

	log.Println("server started at :8080")
	log.Fatal(http.ListenAndServe(":8080",nil))
}
