package prometheus

import "asset/internal/core/port/driven/monitoring"

type prometheus struct {
}

func NewPrometheus() monitoring.Monitoring {
	return prometheus{}
}
