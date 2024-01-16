package logging

type LogContext struct {
	What    string
	When    string
	Where   string
	Why     string
	Who     string
	Message string
}

type Logging interface {
	Trace(context LogContext)
	Debug(context LogContext)
	Info(context LogContext)
	Warn(context LogContext)
	Error(context LogContext)
	Fatal(context LogContext)
}
