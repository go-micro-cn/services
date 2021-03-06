package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"ben/handler"
	"ben/subscriber"

	ben "ben/proto/ben"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.ben"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	ben.RegisterBenHandler(service.Server(), new(handler.Ben))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.ben", service.Server(), new(subscriber.Ben))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
