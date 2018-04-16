package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// StartAPI starts the api server
func StartAPI(bc *Blockchain) {
	r := mux.NewRouter()
	r.HandleFunc("/", helloAPI)

	callbackBC := func(callback func(w http.ResponseWriter, r *http.Request, bc *Blockchain)) func(w http.ResponseWriter, r *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			callback(w, r, bc)
		}
	}
	r.HandleFunc("/block-count", callbackBC(apiGetBlockCount))
	r.HandleFunc("/blocks/{id}", callbackBC(apiGetBlock))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s8", bc.nodeID), r))
}

func helloAPI(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, "hello blockchain")
}
