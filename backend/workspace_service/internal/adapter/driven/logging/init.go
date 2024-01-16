package logging

import (
	// "workspace/configs"
	configuration "workspace/configs"
	loggingDriven "workspace/internal/core/port/driven/logging"

	log "github.com/sirupsen/logrus"
)

type logging struct {
	config configuration.Logging
}

func NewLogger(config configuration.Logging) loggingDriven.Logging {
	return logging{
		config: config,
	}
}

func Init(config configuration.Logging) {

	if config.ENV.ENV == "production" {
		log.SetFormatter(&log.JSONFormatter{})
	} else {
		// The TextFormatter is default, you don't actually have to do this.
		log.SetFormatter(&log.TextFormatter{})
	}
}

//Env
//Format
