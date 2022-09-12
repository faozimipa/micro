module github.com/faozimipa/micro/services.identity

go 1.16

replace github.com/faozimipa/micro/shared => ../shared

require (
	github.com/Nerzal/gocloak/v11 v11.2.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/faozimipa/micro/shared v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.7.7
	github.com/go-playground/validator/v10 v10.10.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.2.0
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/ugorji/go v1.2.7 // indirect
	golang.org/x/sys v0.0.0-20220315194320-039c03cc5b86 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/driver/postgres v1.0.8
	gorm.io/gorm v1.21.3
)
