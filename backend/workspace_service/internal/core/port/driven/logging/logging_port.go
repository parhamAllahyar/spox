package logging

import "time"

type LogContext struct {
	What    string
	When    time.Time
	Where   string
	Why     string
	Who     string
	Message string
}

type Logging interface {
	Trace(LogContext)
	Debug(LogContext)
	Info(LogContext)
	Warn(LogContext)
	Error(LogContext)
	Fatal(LogContext)
}
