package funcs

import (
	"fmt"
	"github.com/LeonZYang/agent/tools/disk"
	"github.com/open-falcon/common/model"
	"log"
)

func DeviceMetrics() (L []*model.MetricValue) {
	mountPoints, err := disk.DiskPartitions(false)
	if err != nil {
		log.Println("Get devices error", err)
		return
	}

	var diskTotal uint64 = 0
	var diskUsed uint64 = 0

	for _, idx := range mountPoints {
		du, err := disk.DiskUsage(idx.Mountpoint)
		if err != nil {
			log.Println("Get device fail: ", err)
			continue
		}
		diskTotal += du.Total
		diskUsed += du.Used
		tags := fmt.Sprintf("mount=%s,fstype=%s", idx.Mountpoint, idx.Fstype)
		L = append(L, GaugeValue("df.bytes.total", du.Total, tags))
		L = append(L, GaugeValue("df.bytes.used", du.Used, tags))
		L = append(L, GaugeValue("df.bytes.free", du.Free, tags))
		L = append(L, GaugeValue("df.bytes.used.percent", du.UsedPercent, tags))
		L = append(L, GaugeValue("df.bytes.free.percent", 100.0-du.UsedPercent, tags))

	}

	if len(L) > 0 && diskTotal > 0 {
		L = append(L, GaugeValue("df.statistics.total", float64(diskTotal)))
		L = append(L, GaugeValue("df.statistics.used", float64(diskUsed)))
		L = append(L, GaugeValue("df.statistics.used.percent", float64(diskUsed)*100.0/float64(diskTotal)))
	}
	return
}
