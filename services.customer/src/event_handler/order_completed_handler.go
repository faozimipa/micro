package eventhandler

import (
	"encoding/json"

	"github.com/faozimipa/micro/services.customer/src/event"
)

func ClearBasket(service *customer.Service, message []byte) {
	var order event.OrderCreated
	json.Unmarshal(message, &order)
	basket, _ := service.GetBasket(order.CustomerID)
	service.ClearBasketItems(basket.ID)
}
