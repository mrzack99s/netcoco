package netautomation

import (
	"math"
	"time"

	"github.com/go-ping/ping"
	"github.com/mrzack99s/netcoco/ent"
)

type PingResponseDeviceMap struct {
	Map     *ent.NetTopologyDeviceMap `json:"map"`
	IsAlive bool                      `json:"is_alive"`
	AvgRtt  string                    `json:"avg_rtt"`
}

type PingResponse struct {
	Map     *ent.Device `json:"map"`
	IsAlive bool        `json:"is_alive"`
	AvgRtt  string      `json:"avg_rtt"`
}

func PingTopology(topo *ent.NetTopology) (response []*PingResponseDeviceMap) {

	for _, obj := range topo.Edges.Topology {
		rep := make(chan *PingResponseDeviceMap)
		go getPingTopology(obj, rep)
		response = append(response, <-rep)
	}

	return
}

func PingDevice(obj *ent.Device) (response *PingResponse) {
	rep := make(chan *PingResponse)
	go getPing(obj, rep)
	response = <-rep

	return
}

func getPing(obj *ent.Device, response chan *PingResponse) {
	pinger, err := ping.NewPinger(obj.DeviceHostname)
	if err != nil {
		panic(err)
	}
	pinger.Count = 1
	pinger.Timeout = time.Second * 1
	pinger.Run()
	stats := pinger.Statistics()
	response <- &PingResponse{
		Map:     obj,
		IsAlive: !(int(math.Ceil(stats.PacketLoss)) > 0),
		AvgRtt:  stats.AvgRtt.String(),
	}
}

func getPingTopology(obj *ent.NetTopologyDeviceMap, response chan *PingResponseDeviceMap) {
	pinger, err := ping.NewPinger(obj.Edges.Device.DeviceHostname)
	if err != nil {
		panic(err)
	}
	pinger.Count = 1
	pinger.SetPrivileged(true)
	pinger.Timeout = time.Second * 1
	pinger.Run()
	stats := pinger.Statistics()
	response <- &PingResponseDeviceMap{
		Map:     obj,
		IsAlive: !(int(math.Ceil(stats.PacketLoss)) > 0),
		AvgRtt:  stats.AvgRtt.String(),
	}
}
