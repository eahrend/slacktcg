package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"

	stcg "github.com/eahrend/slacktcg/mtg"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/mtg", stcg.CardSearch)
	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}
	srv.ListenAndServe()
}
