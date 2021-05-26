package services

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/mrzack99s/netcoco/ent"
	"github.com/mrzack99s/netcoco/ent/device"
	"github.com/mrzack99s/netcoco/ent/nettopology"
	"github.com/mrzack99s/netcoco/pkgs/netautomation"
	"github.com/mrzack99s/netcoco/pkgs/netautomation/types"
	"github.com/mrzack99s/netcoco/pkgs/security"
)

func GetPingTopology(client *ent.Client, id int) (response []*netautomation.PingResponseDeviceMap, err error) {
	topo, e := client.NetTopology.Query().Where(nettopology.IDEQ(id)).Only(context.Background())
	if e != nil {
		err = e
		return
	}

	topo.Edges.Topology = topo.QueryTopology().AllX(context.Background())
	for _, d := range topo.Edges.Topology {
		d.Edges.Device = d.QueryDevice().OnlyX(context.Background())
	}

	response = netautomation.PingTopology(topo)

	return
}

func GetPingDevice(client *ent.Client, id int) (response *netautomation.PingResponse, err error) {
	device, e := client.Device.Query().Where(device.IDEQ(id)).Only(context.Background())
	if e != nil {
		err = e
		return
	}

	response = netautomation.PingDevice(device)

	return
}

func Commit(client *ent.Client, device_id int) (status interface{}, err error) {
	d, e := client.Device.Query().Where(device.IDEQ(device_id)).Only(context.Background())
	d.Edges.InPlatform = d.QueryInPlatform().OnlyX(context.Background())
	d.Edges.Interfaces = d.QueryInterfaces().AllX(context.Background())
	d.Edges.InType = d.QueryInType().OnlyX(context.Background())
	d.Edges.DeletedVlans = d.QueryDeletedVlans().AllX(context.Background())

	if e != nil {
		err = e
		return
	}

	host := types.Host{
		Hostname:     d.DeviceHostname,
		Username:     d.DeviceUsername,
		Password:     security.Decrypt([]byte(d.DevicePassword)),
		EnableSecret: security.Decrypt([]byte(d.DeviceSecret)),
		SSHPort:      d.DeviceSSHPort,
		Platform:     d.Edges.InPlatform.DevicePlatformName,
	}

	task := netautomation.Task{
		Host:     &host,
		Commands: netautomation.GetCommitConfig(d),
	}
	err = task.SendConfig()
	if err != nil {
		status = gin.H{
			"error": err,
		}
		return
	}

	d.Update().SetDeviceCommitConfig(true).Save(context.Background())

	return
}
