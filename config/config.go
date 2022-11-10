package config

import (
	"fmt"
	"gitlab.com/dh-backend/search-service/internal/rabbitMQ"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	GrpcPort        string `mapstructure:"GRPC_PORT" json:"GRPC_PORT"`
	VaultSecretPath string `mapstructure:"VAULT_SECRET_PATH"`
	VaultAddress    string `mapstructure:"VAULT_ADDR"`
	VaultAuthToken  string `mapstructure:"VAULT_AUTH_TOKEN"`
	ConsulAddress   string `mapstructure:"consulAddress" json:"consulAddress"`
	RabbitMQHost    string `mapstructure:"RABBITMQ_HOST" json:"rabbitMQHost"`
	RabbitMQPort    string `mapstructure:"RABBITMQ_PORT" json:"rabbitMQPort"`
	RabbitMQUser    string `mapstructure:"RABBITMQ_USER" json:"rabbitMQUser"`
	RabbitMQPass    string `mapstructure:"RABBITMQ_PASS" json:"rabbitMQPass"`
	CloudAMQPUrl    string `mapstructure:"CLOUDAMQP_URL" json:"cloudAMQPUrl"`
}

func ReadConfigs(path string) *Config {
	viper.AddConfigPath(".")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			viper.AutomaticEnv()

		} else {
			fmt.Printf("cannot read config: %v", err)
		}
	}

	config, err := VaultSecrets(viper.GetString("VAULT_ADDR"), viper.GetString("VAULT_AUTH_TOKEN"), viper.GetString("VAULT_SECRET_PATH"))

	if err != nil {
		log.Fatalf("ERROR: couldn't load secrets: %v", err)
	}

	var rabbitMQUrl string
	if config.CloudAMQPUrl == "" {
		rabbitMQUrl = fmt.Sprintf(
			"amqp://%s:%s@%s:%s",
			config.RabbitMQUser,
			config.RabbitMQPass,
			config.RabbitMQHost,
			config.RabbitMQPort,
		)
	} else {
		rabbitMQUrl = config.CloudAMQPUrl
	}

	// start rabbitMQ connection
	rabbitMQ.ConnectRabbitMq(rabbitMQUrl)

	configs := &Config{
		GrpcPort:      config.GrpcPort,
		ConsulAddress: config.ConsulAddress,
	}

	return configs
}
