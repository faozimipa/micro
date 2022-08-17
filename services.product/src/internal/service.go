package product

import (
	"encoding/json"

	"github.com/faozimipa/micro/services.product/src/entity"
	"github.com/faozimipa/micro/services.product/src/event"
	"github.com/faozimipa/micro/shared/config"
	"github.com/faozimipa/micro/shared/kafka"
	"github.com/google/uuid"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateProduct(product *entity.Product) (*entity.Product, error) {
	product.ID = uuid.New()
	createdProduct, err := s.repo.Create(product)

	if err != nil {
		return createdProduct, err
	}

	event := event.ProcuctCreated{
		ID:       createdProduct.ID,
		Name:     createdProduct.Name,
		Price:    createdProduct.Price,
		Quantity: createdProduct.Quantity,
	}
	payload, _ := json.Marshal(event)
	kafka.Publish(product.ID, payload, "ProductCreated", config.AppConfig.KafkaProductTopic)

	return createdProduct, nil
}

func (s *Service) BulkUpdate(products *[]entity.Product) error {
	return s.repo.BulkUpdate(products)
}

func (s *Service) GetProducts() ([]entity.Product, error) {
	return s.repo.GetList()
}

func (s *Service) GetProduct(id uuid.UUID) (entity.Product, error) {
	return s.repo.GetById(id)
}
