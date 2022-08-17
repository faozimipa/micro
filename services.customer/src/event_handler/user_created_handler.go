package eventhandler

import (
	"encoding/json"

	"github.com/faozimipa/micro/services.customer/src/entity"
)

func CreateCustomer(service *customer.Service, message []byte) {
	var customer entity.Customer
	json.Unmarshal(message, &customer)
	service.CreateCustomer(customer)
}
