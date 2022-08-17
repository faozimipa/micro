module github.com/faozimipa/micro/services.customer

go 1.16

replace github.com/faozimipa/micro/shared => ../shared

require (
	github.com/gin-gonic/gin v1.8.1
	github.com/google/uuid v1.3.0
	github.com/pact-foundation/pact-go v1.5.1
	github.com/pelletier/go-toml/v2 v2.0.3 // indirect
	github.com/satori/go.uuid v1.2.0
	github.com/spf13/afero v1.9.2 // indirect
	github.com/spf13/viper v1.12.0 // indirect
	github.com/faozimipa/micro/shared v0.0.0-00010101000000-000000000000
	github.com/subosito/gotenv v1.4.0 // indirect
	golang.org/x/sys v0.0.0-20220817070843-5a390386f1f2 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gorm.io/driver/postgres v1.0.8
	gorm.io/gorm v1.21.3
)
