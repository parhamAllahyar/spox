package configs

import "os"

type PostgresDBConfig struct {
	User         string `mapstructure:"POSTGRES_USER"`
	Password     string `mapstructure:"POSTGRES_PASSWORD"`
	Host         string `mapstructure:"POSTGRES_HOST"`
	Port         string `mapstructure:"POSTGRES_PORT"`
	DatabaseName string `mapstructure:"POSTGRES_DB"`
}

type HttpServer struct {
	Host string `mapstructure:"ADMIN_APP_HTTP_HOST"`
	Port string `mapstructure:"ADMIN_APP_HTTP_PORT"`
}

type Config struct {
	Postgre    PostgresDBConfig `mapstructure:"postgre"`
	HttpServer HttpServer       `mapstructure:"httpServer"`
	ENV        ENV              `mapstructure:"env"`
}

type Logging struct {
	Format string
	ENV    ENV
}

type ENV struct {
	ENV string
}

func LoadConfig() Config {

	return Config{
		PostgresDBConfig{
			User:         os.Getenv("POSTGRES_USER"),
			Password:     os.Getenv("POSTGRES_PASSWORD"),
			Host:         os.Getenv("POSTGRES_HOST"),
			Port:         os.Getenv("POSTGRES_PORT"),
			DatabaseName: os.Getenv("POSTGRES_DB"),
		},

		HttpServer{
			Port: os.Getenv("ADMIN_APP_HTTP_PORT"),
			Host: os.Getenv("ADMIN_APP_HTTP_HOST"),
		},

		ENV{
			ENV: os.Getenv("ENV"),
		},
	}

}
