package identity

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Nerzal/gocloak/v11"
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

	IsUsernameExist := s.repo.IsUsernameExist(user.Username)
	if IsUsernameExist {
		return user, errors.New("Username is already exist!")
	}

	isEmailExist := s.repo.IsUserExist(user.Email)
	if isEmailExist {
		return user, errors.New("Email is already exist!")
	}

	var url = fmt.Sprintf("http://%s:%s", config.AppConfig.KeyloackHost, config.AppConfig.KeyloackPort)
	client := gocloak.NewClient(url)
	ctx := context.Background()
	token, err := client.LoginAdmin(ctx, config.AppConfig.KeyloackUsername, config.AppConfig.KeyloackPassword, config.AppConfig.KeyloackRealm)
	if err != nil {
		fmt.Println(err.Error())
		return user, errors.New("Something wrong on login!")
	}
   
	userKeyloack := gocloak.User{
	 FirstName: 	gocloak.StringP(user.FirstName),
	 LastName:  	gocloak.StringP(user.LastName),
	 Email:     	gocloak.StringP(user.Email),
	 Enabled:   	gocloak.BoolP(true),
	 Username:  	gocloak.StringP(user.Username),
	}
   
	userIDKeyLoack, err := client.CreateUser(ctx, token.AccessToken, config.AppConfig.KeyloackRealm, userKeyloack)
	if err != nil {
		return user, errors.New("Something wrong on create!")
	} else {
		fmt.Println("user created wwith id :")
		fmt.Println(userIDKeyLoack)
		uid, _ := uuid.Parse(userIDKeyLoack)
		user.ID = uid
	}
	
	usr, err := s.repo.Create(user)
	if err != nil {
		return user, err
	}

	event := event.UserCreated{
		ID:        usr.ID,
		Email:     usr.Email,
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
		Username:  usr.Username,
	}

	payload, _ := json.Marshal(event)
	kafka.Publish(user.ID, payload, "UserCreated", config.AppConfig.KafkaUserTopic)

	return user, nil
}

func (s *Service) ValidateUser(email string, password string) (entity.User, error) {
	return s.repo.GetUserByEmail(email, password)
}

func (s *Service) Login(username string, password string) (*gocloak.JWT, error) {
	var url = fmt.Sprintf("http://%s:%s", config.AppConfig.KeyloackHost, config.AppConfig.KeyloackPort)
	client := gocloak.NewClient(url)
	ctx := context.Background()
	token, err := client.Login(ctx, config.AppConfig.KeyloackClientId, config.AppConfig.KeyloackClientSecret, config.AppConfig.KeyloackRealm, username, password)
	if err != nil {
		fmt.Println(err.Error())
		return nil, errors.New("Invalid Credentials.")
	}
	return token, nil
}