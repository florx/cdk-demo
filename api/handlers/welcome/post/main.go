package main

import (
	"github.com/akrylysov/algnhsa"
	"github.com/florx/cdk-demo/api/handlers/welcome"
	"go.uber.org/zap"
)

func main() {
	log, err := zap.NewProduction()
	if err != nil {
		panic("failed to create logger: " + err.Error())
	}

	welcomeHandler := welcome.NewHandler(log)
	algnhsa.ListenAndServe(welcomeHandler, nil)
}
