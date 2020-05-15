package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func storeState(r *http.Request) {
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	state = body
}

var state []byte

func handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		log.Println("GET")
		w.WriteHeader(http.StatusOK)
		w.Write(state)
	case http.MethodDelete:
		log.Println("DELETE")
		state = []byte("")
		w.WriteHeader(http.StatusGone)
	case http.MethodPost:
		log.Println("POST")
		storeState(r)
		w.WriteHeader(http.StatusCreated)
	}
}

func main() {
	http.HandleFunc("/", handle)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("start sever")
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal("failed to start server")
	}
}
