package monitor

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"strings"
	"time"
)

type Monitor struct {
	completed   prometheus.Counter
	reqError    prometheus.Counter
	reqDuration prometheus.Summary
}

func MethodLabel(method string) map[string]string {
	return map[string]string{
		"method": strings.ToLower(method),
	}
}

func NewMonitor(name string) *Monitor {
	quantiles := map[float64]float64{
		0.5:  0.05,
		0.8:  0.01,
		0.9:  0.01,
		0.95: 0.001,
		0.99: 0.001,
	}
	monitor := Monitor{
		completed: promauto.NewCounter(prometheus.CounterOpts{
			Name: strings.ToLower(name) + CompleteTotal,
		}),
		reqError: promauto.NewCounter(prometheus.CounterOpts{
			Name: strings.ToLower(name) + RequestErrorTotal,
		}),
		reqDuration: promauto.NewSummary(prometheus.SummaryOpts{
			Name:       strings.ToLower(name) + RequestDuration,
			Objectives: quantiles,
			MaxAge:     time.Minute,
		}),
	}
	return &monitor

}

func (m *Monitor) IncComplete() {
	m.completed.Inc()
}
func (m *Monitor) IncReqError() {
	m.reqError.Inc()
}
