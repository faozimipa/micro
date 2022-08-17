package main

import (
	"github.com/faozimipa/micro/services.notification/src/kafka"
	"github.com/faozimipa/micro/shared/config"
)

func main() {
	config := config.LoadConfig(".")
	kafka.RegisterConsumer(config.KafkaOrderTopic)
}
