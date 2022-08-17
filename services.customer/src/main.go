package main

import (
	"github.com/faozimipa/micro/services.customer/src/entity"
	customer "github.com/faozimipa/micro/services.customer/src/internal"
	"github.com/faozimipa/micro/services.customer/src/kafka"
	"github.com/faozimipa/micro/shared/config"
	"github.com/faozimipa/micro/shared/server"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	config := config.LoadConfig(".")

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  config.GetDBURL(),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		panic("Couldn't connect to the DB.")
	}

	db.AutoMigrate(&entity.Basket{})
	db.AutoMigrate(&entity.BasketItem{})
	db.AutoMigrate(&entity.Customer{})
	db.AutoMigrate(&entity.Product{})

	repo := customer.NewRepository(db)
	service := customer.NewService(repo)
	handler := customer.NewHandler(service)

	go kafka.RegisterConsumer(config.KafkaUserTopic, service)
	go kafka.RegisterConsumer(config.KafkaProductTopic, service)
	go kafka.RegisterConsumer(config.KafkaOrderTopic, service)

	err = server.NewServer(handler.Init(), config.AppPort).Run()
	if err != nil {
		panic("Couldn't start the HTTP server.")
	}
}
