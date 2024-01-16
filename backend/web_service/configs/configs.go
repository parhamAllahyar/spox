package configs

import (
	"fmt"
	"os"
)

type Services struct {
	UserPort      string
	UserHost      string
	WorkspaceHost string
	WorkspacePORT string
	BoardPort     string
	BoardHost     string
	AssetPort     string
	AssetHost     string
}

type HttpServer struct {
	Host string `mapstructure:"HOST"`
	Port string `mapstructure:"PORT"`
}

type Configs struct {
	HTTPServer HttpServer
	Services   Services
}

func LoadConfig() Configs {

	fmt.Println("port", os.Getenv("USER_APP_PORT"))

	fmt.Println("host", os.Getenv("USER_APP_HOST"))

	return Configs{
		HttpServer{
			Port: os.Getenv("PORT"),
			Host: os.Getenv(""),
		},
		Services{
			UserPort:      os.Getenv("USER_APP_PORT"),
			UserHost:      os.Getenv("USER_APP_HOST"),
			WorkspaceHost: os.Getenv(""),
			WorkspacePORT: os.Getenv(""),
			BoardPort:     os.Getenv(""),
			BoardHost:     os.Getenv(""),
			AssetPort:     os.Getenv(""),
			AssetHost:     os.Getenv(""),
		},
	}
}
