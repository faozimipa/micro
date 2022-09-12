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

	//get token from user keyloack
	// tokenKey, err:= keyloack.GetToken()
	// if err != nil{
	// 	return user, errors.New("Something wrong!")
	// }

	client := gocloak.NewClient(config.AppConfig.KeyloackHost)
	ctx := context.Background()
	token, err := client.LoginAdmin(ctx, config.AppConfig.KeyloackUsername, config.AppConfig.KeyloackPassword, config.AppConfig.KeyloackRealm)
	if err != nil {
		return user, errors.New("Something wrong!")
	}
   
	userKeyloack := gocloak.User{
	 FirstName: 	gocloak.StringP(user.FirstName),
	 LastName:  	gocloak.StringP(user.LastName),
	 Email:     	gocloak.StringP(user.Email),
	 Enabled:   	gocloak.BoolP(true),
	 Username:  	gocloak.StringP(user.Username),
	//  Credentials: 	gocloak.User.Credentials{
	// 	Temporary: 	gocloak.BoolP(false),
	// 	Type:		gocloak.StringP("password"),
	// 	Value: 		gocloak.StringP(user.Password),
	//  },
	}
   
	userIDKeyLoack, err := client.CreateUser(ctx, token.AccessToken, config.AppConfig.KeyloackRealm, userKeyloack)
	if err != nil {
		return user, errors.New("Something wrong!")
	} else {
		fmt.Println("user created wwith id :")
		fmt.Println(userIDKeyLoack)
		uid, _ := uuid.Parse(userIDKeyLoack)
		user.ID = uid
	}

	// registeredUser, err := keyloack.RegisterUser(user, tokenKey)
	// if err != nil{
	// 	return user, errors.New("Something wrong!")
	// }else{
	// 	user.ID = registeredUser.ID
	// }


	// user.ID = uuid.New()
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
