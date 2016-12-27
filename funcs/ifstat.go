package funcs

import (
	"github.com/LeonZYang/agent/tools/net"
	"github.com/open-falcon/common/model"
	"log"
)

func NetMetrics() []*model.MetricValue {
	return CoreNetMetrics()
}

func CoreNetMetrics() []*model.MetricValue {

	netIfs, err := net.NetIOCounters(true)
	if err != nil {
		log.Println("Get netInfo fail: ", err)
		return []*model.MetricValue{}
	}

	cnt := len(netIfs)
	ret := make([]*model.MetricValue, cnt*8)

	for idx, netIf := range netIfs {
		iface := "iface=" + netIf.Name
		ret[idx*8+0] = CounterValue("net.if.in.bytes", netIf.BytesRecv, iface)
		ret[idx*8+1] = CounterValue("net.if.in.packets", netIf.PacketsRecv, iface)
		ret[idx*8+2] = CounterValue("net.if.in.errors", netIf.Errin, iface)
		ret[idx*8+3] = CounterValue("net.if.in.dropped", netIf.Dropin, iface)
		ret[idx*8+4] = CounterValue("net.if.out.bytes", netIf.BytesSent, iface)
		ret[idx*8+5] = CounterValue("net.if.out.packets", netIf.PacketsSent, iface)
		ret[idx*8+6] = CounterValue("net.if.out.errors", netIf.Errout, iface)
		ret[idx*8+7] = CounterValue("net.if.out.dropped", netIf.Dropout, iface)

	}
	return ret
}
