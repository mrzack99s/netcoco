package services

import (
	"context"
	"errors"
	"log"

	"github.com/mrzack99s/netcoco/ent"
	"github.com/mrzack99s/netcoco/ent/device"
	"github.com/mrzack99s/netcoco/ent/devicetype"
	"github.com/mrzack99s/netcoco/ent/ipstaticroutingtable"
	"github.com/mrzack99s/netcoco/ent/netinterface"
	"github.com/mrzack99s/netcoco/ent/nettopologydevicemap"
)

func GetCanBeStaticRouteOnDevice(client *ent.Client, did int) (response []*ent.Device, err error) {
	deviceConnected, err := client.NetTopologyDeviceMap.Query().Where(
		nettopologydevicemap.And(
			nettopologydevicemap.HasEdgeWith(
				nettopologydevicemap.HasDeviceWith(device.IDEQ(did)),
			),
			nettopologydevicemap.HasDeviceWith(
				device.HasInTypeWith(
					devicetype.Or(
						devicetype.DeviceTypeNameEQ("router"),
						devicetype.DeviceTypeNameEQ("l3switch"),
					),
				),
			),
		),
	).
		WithDevice().
		All(context.Background())
	if err != nil {
		return
	}

	if len(deviceConnected) > 0 {
		for _, d := range deviceConnected {
			d.Edges.Device.Edges.HaveIPAddresses = d.Edges.Device.QueryHaveIPAddresses().AllX(context.Background())
			response = append(response, d.Edges.Device)
		}
	} else {
		err = errors.New("cannot be route to another device")
	}

	return
}

func CreateIPRouting(client *ent.Client, obj ent.IPStaticRoutingTable) (response *ent.IPStaticRoutingTable, err error) {
	if i := client.IPStaticRoutingTable.Query().Where(
		ipstaticroutingtable.And(
			ipstaticroutingtable.NetworkAddressEQ(obj.NetworkAddress),
			ipstaticroutingtable.SubnetMaskEQ(obj.SubnetMask),
			ipstaticroutingtable.NextHopEQ(obj.NextHop),
			ipstaticroutingtable.HasOnDeviceWith(device.IDEQ(obj.Edges.OnDevice.ID)),
			ipstaticroutingtable.HasOnInterfaceWith(netinterface.IDEQ(obj.Edges.OnInterface.ID)),
		),
	).CountX(context.Background()); i > 0 {
		err = errors.New("ip routing duplicated")
		return
	}
	response, err = client.IPStaticRoutingTable.Create().
		SetNetworkAddress(obj.NetworkAddress).
		SetSubnetMask(obj.SubnetMask).
		SetNextHop(obj.NextHop).
		SetBrdInterface(obj.BrdInterface).
		SetOnInterface(obj.Edges.OnInterface).
		SetOnDevice(obj.Edges.OnDevice).
		Save(context.Background())

	response.Edges.OnDevice = response.QueryOnDevice().OnlyX(context.Background())
	if response.Edges.OnDevice.DeviceCommitConfig {
		_, err := response.Edges.OnDevice.Update().SetDeviceCommitConfig(false).Save(context.Background())
		if err != nil {
			log.Fatal(err)
		}
	}

	return
}

func DeleteIPRouting(client *ent.Client, id int) (err error) {
	d, err := client.IPStaticRoutingTable.Query().Where(ipstaticroutingtable.IDEQ(id)).WithOnDevice().Only(context.Background())
	if err != nil {
		return
	}

	err = client.IPStaticRoutingTable.DeleteOneID(d.ID).Exec(context.Background())
	if d.Edges.OnDevice.DeviceCommitConfig {
		_, err := d.Edges.OnDevice.Update().SetDeviceCommitConfig(false).Save(context.Background())
		if err != nil {
			log.Fatal(err)
		}
	}
	return
}
