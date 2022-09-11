module github.com/faozimipa/micro/services.identity

go 1.16

replace github.com/faozimipa/micro/shared => ../shared

require (
	github.com/Nerzal/gocloak/v11 v11.2.0 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/faozimipa/micro/shared v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.7.7
	github.com/google/uuid v1.2.0
	github.com/tbaehler/gin-keycloak v1.3.0 // indirect
	gorm.io/driver/postgres v1.0.8
	gorm.io/gorm v1.21.3
)
