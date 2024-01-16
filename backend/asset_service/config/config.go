package config

import "os"

type S3 struct {
	Credentials string `mapstructure:"CREDENTIALS"`
	Region      string `mapstructure:"REGION"`
	Endpoint    string `mapstructure:"ENDPOINT"`
	Bucket      string `mapstructure:"BUCKET"`
	ACCESSKEY   string `mapstructure:"ACCESSKEY"`
	SECRETKEY   string `mapstructure:"SECRETKEY"`
	Token       string `mapstructure:"TOKEN"`
}

type PostgresDBConfig struct {
	User         string `mapstructure:"POSTGRES_USER"`
	Password     string `mapstructure:"POSTGRES_PASSWORD"`
	Host         string `mapstructure:"POSTGRES_URL"`
	Port         string `mapstructure:"POSTGRES_PORT"`
	DatabaseName string `mapstructure:"POSTGRES_DB"`
}

type HttpServer struct {
	Host string `mapstructure:"USER_APP_HOST"`
	Port string `mapstructure:"USER_APP_PORT"`
}

type Token struct {
	User_Token  string `mapstructure:"USER_TOKEN"`
	Admin_Token string `mapstructure:"ADMIN_TOKEN"`
}

type Config struct {
	Postgre    PostgresDBConfig
	HttpServer HttpServer
	ENV        ENV
	Token      Token
	S3         S3
}

type ENV struct {
	ENV string `mapstructure:"env"`
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
			Port: os.Getenv("USER_HTTP_PORT"),
			Host: os.Getenv("USER_HTTP_HOST"),
		},

		ENV{
			ENV: os.Getenv("ENV"),
		},

		Token{
			User_Token:  os.Getenv("USER_TOKEN"),
			Admin_Token: os.Getenv("ADMIN_TOKEN"),
		},

		S3{
			Credentials: os.Getenv("CREDENTIALS"),
			Region:      os.Getenv("REGION"),
			Endpoint:    os.Getenv("ENDPOINT"),
			Bucket:      os.Getenv("BUCKET"),
			ACCESSKEY:   os.Getenv("ACCESSKEY"),
			SECRETKEY:   os.Getenv("SECRETKEY"),
			Token:       os.Getenv("TOKEN"),
		},
	}

}
