package eventhandler

import (
	"encoding/json"

	"github.com/faozimipa/micro/services.customer/src/entity"
	"github.com/faozimipa/micro/services.customer/src/internal"
)

func CreateProduct(service *customer.Service, message []byte) {
	var product entity.Product
	json.Unmarshal(message, &product)
	service.CreateProduct(product)
}
