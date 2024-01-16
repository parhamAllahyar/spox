package logging

import (
	loggingDriven "workspace/internal/core/port/driven/logging"

	log "github.com/sirupsen/logrus"
)

func (l logging) Trace(context loggingDriven.LogContext) {
	log.WithFields(log.Fields{
		"what":  context.What,
		"When":  context.When,
		"Where": context.Where,
		"Why":   context.Why,
		"Who":   context.Who,
	}).Fatal(context.Message)
}

func (l logging) Debug(context loggingDriven.LogContext) {

	log.WithFields(log.Fields{
		"what":  context.What,
		"When":  context.When,
		"Where": context.Where,
		"Why":   context.Why,
		"Who":   context.Who,
	}).Fatal(context.Message)

}

func (l logging) Info(context loggingDriven.LogContext) {

	log.WithFields(log.Fields{
		"what":  context.What,
		"When":  context.When,
		"Where": context.Where,
		"Why":   context.Why,
		"Who":   context.Who,
	}).Fatal(context.Message)

}

func (l logging) Warn(context loggingDriven.LogContext) {

	log.WithFields(log.Fields{
		"what":  context.What,
		"When":  context.When,
		"Where": context.Where,
		"Why":   context.Why,
		"Who":   context.Who,
	}).Fatal(context.Message)

}

func (l logging) Error(context loggingDriven.LogContext) {

	log.WithFields(log.Fields{
		"what":  context.What,
		"When":  context.When,
		"Where": context.Where,
		"Why":   context.Why,
		"Who":   context.Who,
	}).Fatal(context.Message)

}

func (l logging) Fatal(context loggingDriven.LogContext) {
	log.WithFields(log.Fields{
		"what":  context.What,
		"When":  context.When,
		"Where": context.Where,
		"Why":   context.Why,
		"Who":   context.Who,
	}).Fatal(context.Message)
}
