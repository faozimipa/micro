package identity

import (
	"encoding/json"
	"errors"

	"github.com/faozimipa/micro/services.identity/src/entity"
	"github.com/faozimipa/micro/services.identity/src/event"
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

func (s *Service) SignUp(user *entity.User) (*entity.User, error) {

	isEmailExist := s.repo.IsUserExist(user.Email)
	if isEmailExist {
		return user, errors.New("Email is already exist!")
	}

	user.ID = uuid.New()
	usr, err := s.repo.Create(user)
	if err != nil {
		return user, err
	}

	event := event.UserCreated{
		ID:        usr.ID,
		Email:     usr.Email,
		FirstName: usr.FirstName,
		LastName:  usr.LastName}

	payload, _ := json.Marshal(event)
	kafka.Publish(user.ID, payload, "UserCreated", config.AppConfig.KafkaUserTopic)

	return user, nil
}

func (s *Service) ValidateUser(email string, password string) (entity.User, error) {
	return s.repo.GetUserByEmail(email, password)
}
