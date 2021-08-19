package main

import (
	"net/http"

	"github.com/florx/cdk-demo/api/handlers/welcome"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func main() {

	log, _ := zap.NewProduction()

	welcomeHandler := welcome.NewHandler(log)

	r := mux.NewRouter()
	r.HandleFunc("/", welcomeHandler.Post)
	http.ListenAndServe(":8000", r)
}
