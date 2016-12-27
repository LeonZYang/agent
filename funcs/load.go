package funcs

import (
	"github.com/LeonZYang/agent/tools/load"
	"github.com/open-falcon/common/model"
	"log"
)

func LoadMetrics() (L []*model.MetricValue) {

	loadVal, err := load.LoadAvg()
	if err != nil {
		log.Println("Get load fail: ", err)
		return nil
	}

	L = append(L, CounterValue("load.load1", loadVal.Load1))
	L = append(L, CounterValue("load.load5", loadVal.Load5))
	L = append(L, CounterValue("load.load15", loadVal.Load15))

	return
}
