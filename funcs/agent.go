package funcs

import (
	"github.com/open-falcon/common/model"
)

func AgentMetrics() (L []*model.MetricValue) {
	L = append(L, GaugeValue("agent.alive", 1))
	return
}
