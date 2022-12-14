package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var AppConfig Config

type Config struct {
	DbDriver   string `mapstructure:"DB_DRIVER"`
	DbUser     string `mapstructure:"DB_USER"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbPort     string `mapstructure:"DB_PORT"`
	DbHost     string `mapstructure:"DB_HOST"`
	DbName     string `mapstructure:"DB_NAME"`
	AppPort    string `mapstructure:"APP_PORT"`

	KeyloackHost         string `mapstructure:"KEY_HOST"`
	KeyloackPort         string `mapstructure:"KEY_PORT"`
	KeyloackRealm        string `mapstructure:"KEY_REALM"`
	KeyloackUsername     string `mapstructure:"KEY_USERNAME"`
	KeyloackPassword     string `mapstructure:"KEY_PASSWORD"`
	KeyloackGrantType    string `mapstructure:"KEY_GRANT_TYPE"`
	KeyloackClientId     string `mapstructure:"KEY_CLIENT_ID"`
	KeyloackClientSecret string `mapstructure:"KEY_CLIENT_SECRET"`

	KafkaUserTopic          string `mapstructure:"KAFKA_USER_TOPIC"`
	KafkaProductTopic       string `mapstructure:"KAFKA_PRODUCT_TOPIC"`
	KafkaOrderTopic         string `mapstructure:"KAFKA_ORDER_TOPIC"`
	KafkaBrokerAddress      string `mapstructure:"KAFKA_BROKER_ADDRESS"`
	CustomerServiceEndpoint string `mapstructure:"CUSTOMER_SERVICE_ENDPOINT"`
}

func (c Config) getKeyloackUrl() string {
	return fmt.Sprintf("http://localhost:8080/auth/realms/%s/protocol/openid-connect", c.KeyloackRealm)
}

func (c Config) GetDBURL() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		c.DbHost, c.DbPort, c.DbUser, c.DbName, c.DbPassword)
}

func LoadConfig(path string) (config Config) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	viper.ReadInConfig()
	viper.Unmarshal(&config)
	AppConfig = config
	return AppConfig
}
