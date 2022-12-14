package eventhandler

import (
	"encoding/json"
	"log"

	"github.com/faozimipa/micro/services.order/src/entity"
	order "github.com/faozimipa/micro/services.order/src/internal"
	"github.com/faozimipa/micro/shared/config"
	"github.com/faozimipa/micro/shared/kafka"
	"github.com/google/uuid"
)

func FailOrder(service *order.Service, messageKey []byte) {

	orderID, _ := uuid.Parse(string(messageKey))
	order, err := service.UpdateOrderStatus(uuid.UUID(orderID), int(entity.OrderFailed))
	if err != nil {
		log.Printf("FailOrder.UpdateOrderStatus failed: %v", err.Error())
		return
	}
	payload, _ := json.Marshal(order)
	kafka.Publish(order.ID, payload, "OrderFailed", config.AppConfig.KafkaOrderTopic)
}
