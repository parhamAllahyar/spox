package configs

import "os"

type CassandaraDBConfig struct {
	User         string `mapstructure:"POSTGRES_USER"`
	Password     string `mapstructure:"POSTGRES_PASSWORD"`
	Host         string `mapstructure:"POSTGRES_URL"`
	Port         string `mapstructure:"POSTGRES_PORT"`
	DatabaseName string `mapstructure:"POSTGRES_DB"`
}

type HttpServer struct {
	Host string `mapstructure:"BOARD_APP_GRPC_HOST"`
	Port string `mapstructure:"BOARD_APP_GRPC_PORT"`
}

type Config struct {
	CassandaraDBConfig CassandaraDBConfig `mapstructure:"postgre"`
	HttpServer         HttpServer         `mapstructure:"grpcServer"`
}

func LoadConfig() Config {

	//TODO: use viper

	return Config{
		CassandaraDBConfig{
			User:         os.Getenv("POSTGRES_USER"),
			Password:     os.Getenv("POSTGRES_PASSWORD"),
			Host:         os.Getenv("POSTGRES_URL"),
			Port:         os.Getenv("POSTGRES_PORT"),
			DatabaseName: os.Getenv("POSTGRES_DB"),
		},

		HttpServer{
			Port: os.Getenv("BOARD_APP_GRPC_HOST"),
			Host: os.Getenv("BOARD_APP_GRPC_PORT"),
		},
	}

}
