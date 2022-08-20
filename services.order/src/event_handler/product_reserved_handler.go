package eventhandler

import (
	"encoding/json"
	"log"

	"github.com/faozimipa/micro/services.order/src/entity"
	"github.com/faozimipa/micro/services.order/src/internal"
	"github.com/faozimipa/micro/shared/config"
	"github.com/faozimipa/micro/shared/kafka"
	"github.com/google/uuid"
	gouuid "github.com/satori/go.uuid"
)

func CompleteOrder(service *order.Service, messageKey []byte) {

	orderID, _ := gouuid.FromString(string(messageKey))
	order, err := service.UpdateOrderStatus(uuid.UUID(orderID), int(entity.OrderCompleted))
	if err != nil {
		log.Printf("CompleteOrder.UpdateOrderStatus failed: %v", err.Error())
		return
	}
	payload, _ := json.Marshal(order)
	kafka.Publish(order.ID, payload, "OrderCompleted", config.AppConfig.KafkaOrderTopic)
}
