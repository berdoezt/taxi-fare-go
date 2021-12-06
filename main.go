package main

import (
	"github.com/berdoezt/taxi-fare-go/app/handler"
	"github.com/berdoezt/taxi-fare-go/app/service"
	"github.com/berdoezt/taxi-fare-go/app/validation"
	"github.com/berdoezt/taxi-fare-go/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	config.InitLogrus(cfg)

	fareService := service.NewFareServiceImpl(
		service.WithFareRules(cfg.FareRules),
	)
	fareValidation := validation.NewFareValidation()
	fareHandler := handler.NewFareHandler(
		handler.WithValidator(fareValidation),
		handler.WithService(fareService),
	)

	fareHandler.Handle()
}
