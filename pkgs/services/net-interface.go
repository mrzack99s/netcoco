package services

import (
	"context"
	"errors"
	"log"

	"github.com/mrzack99s/netcoco/ent"
	"github.com/mrzack99s/netcoco/ent/device"
	"github.com/mrzack99s/netcoco/ent/ipaddress"
	"github.com/mrzack99s/netcoco/ent/netinterface"
	"github.com/mrzack99s/netcoco/ent/netinterfacelayer"
	"github.com/mrzack99s/netcoco/ent/netinterfacemode"
	"github.com/mrzack99s/netcoco/ent/vlan"
)

func GetInterface(client *ent.Client, id int) (response *ent.NetInterface, err error) {
	response, err = client.NetInterface.Query().Where(netinterface.IDEQ(id)).Only(context.Background())
	response.Edges.Mode = response.QueryMode().OnlyX(context.Background())
	return
}

func CreateInterface(client *ent.Client, obj ent.NetInterface) (response *ent.NetInterface, err error) {

	switch obj.Edges.OnLayer.InterfaceLayer {
	case 2:
		if obj.Edges.Mode.InterfaceMode == "EtherChannel" {
			response, err = client.NetInterface.Create().
				SetInterfaceName(obj.InterfaceName).
				SetInterfaceShutdown(obj.InterfaceShutdown).
				SetOnDevice(obj.Edges.OnDevice).
				SetOnLayer(obj.Edges.OnLayer).
				SetOnPoInterface(obj.Edges.OnPoInterface).
				SetMode(obj.Edges.Mode).
				Save(context.Background())

		} else if obj.Edges.NativeOnVlan == nil {
			response, err = client.NetInterface.Create().
				SetInterfaceName(obj.InterfaceName).
				SetInterfaceShutdown(obj.InterfaceShutdown).
				SetOnDevice(obj.Edges.OnDevice).
				SetOnLayer(obj.Edges.OnLayer).
				AddHaveVlans(obj.Edges.HaveVlans...).
				SetMode(obj.Edges.Mode).
				Save(context.Background())

		} else if obj.Edges.HaveVlans == nil {
			response, err = client.NetInterface.Create().
				SetInterfaceName(obj.InterfaceName).
				SetInterfaceShutdown(obj.InterfaceShutdown).
				SetOnDevice(obj.Edges.OnDevice).
				SetNativeOnVlan(obj.Edges.NativeOnVlan).
				SetMode(obj.Edges.Mode).
				SetOnLayer(obj.Edges.OnLayer).
				Save(context.Background())

		} else if obj.Edges.NativeOnVlan == nil && obj.Edges.HaveVlans == nil {
			response, err = client.NetInterface.Create().
				SetInterfaceName(obj.InterfaceName).
				SetInterfaceShutdown(obj.InterfaceShutdown).
				SetOnDevice(obj.Edges.OnDevice).
				SetMode(obj.Edges.Mode).
				SetOnLayer(obj.Edges.OnLayer).
				Save(context.Background())

		} else {
			response, err = client.NetInterface.Create().
				SetInterfaceName(obj.InterfaceName).
				SetInterfaceShutdown(obj.InterfaceShutdown).
				SetOnDevice(obj.Edges.OnDevice).
				SetNativeOnVlan(obj.Edges.NativeOnVlan).
				SetMode(obj.Edges.Mode).
				AddHaveVlans(obj.Edges.HaveVlans...).
				SetOnLayer(obj.Edges.OnLayer).
				Save(context.Background())

		}

	case 3:
		ipAddr, e := CheckNotDuplicateIPAddress(client, obj.Edges.OnDevice, obj.Edges.OnIPAddress)
		if e != nil {
			ipAddr = client.IPAddress.Create().
				SetIPAddress(obj.Edges.OnIPAddress.IPAddress).
				SetSubnetMask(obj.Edges.OnIPAddress.SubnetMask).
				SetOnDevice(obj.Edges.OnDevice).
				SaveX(context.Background())
		} else {
			if ipAddr.ID != obj.Edges.OnIPAddress.ID {
				err = errors.New("ip duplicated")
				return
			}
		}

		modeNone := client.NetInterfaceMode.Query().Where(netinterfacemode.InterfaceModeEQ("None")).OnlyX(context.Background())
		response, err = client.NetInterface.Create().
			SetInterfaceName(obj.InterfaceName).
			SetInterfaceShutdown(obj.InterfaceShutdown).
			SetOnDevice(obj.Edges.OnDevice).
			SetMode(modeNone).
			SetOnIPAddress(ipAddr).
			SetOnLayer(obj.Edges.OnLayer).
			Save(context.Background())

	}

	if obj.Edges.OnDevice.DeviceCommitConfig {
		_, err = obj.Edges.OnDevice.Update().SetDeviceCommitConfig(false).Save(context.Background())
		if err != nil {
			log.Fatal(err)
		}
	}

	return
}

func CreateRangeInterface(client *ent.Client, obj []ent.NetInterface) (response []*ent.NetInterface, err error) {
	vlan1, err := client.Vlan.Query().Where(vlan.VlanIDEQ(1)).Only(context.Background())
	if err != nil {
		vlan1, _ = client.Vlan.Create().SetVlanID(1).Save(context.Background())
	}

	layer := client.NetInterfaceLayer.Query().Where(netinterfacelayer.InterfaceLayerEQ(2)).OnlyX(context.Background())
	device := client.Device.Query().Where(device.IDEQ(obj[0].Edges.OnDevice.ID)).
		WithInType().
		OnlyX(context.Background())

	for _, item := range obj {
		if device.Edges.InType.DeviceTypeName == "router" {
			layer = client.NetInterfaceLayer.Query().Where(netinterfacelayer.InterfaceLayerEQ(3)).OnlyX(context.Background())
		}

		r, e := client.NetInterface.Create().
			SetInterfaceName(item.InterfaceName).
			SetOnDevice(item.Edges.OnDevice).
			SetNativeOnVlan(vlan1).
			SetMode(item.Edges.Mode).
			SetOnLayer(layer).
			AddHaveVlans(vlan1).
			Save(context.Background())
		if e != nil {
			err = e
			return
		} else {
			response = append(response, r)
		}
	}

	if device.DeviceCommitConfig {
		_, err = device.Update().SetDeviceCommitConfig(false).Save(context.Background())
		if err != nil {
			log.Fatal(err)
		}
	}

	return
}

func EditInterfaceDetail(client *ent.Client, obj ent.NetInterface) (response *ent.NetInterface, err error) {
	usr, err := client.NetInterface.
		Query().Where(netinterface.IDEQ(obj.ID)).
		Only(context.Background())
	usr.Edges.OnDevice = usr.QueryOnDevice().OnlyX(context.Background())
	if err != nil {
		response = nil
		err = errors.New("not found interface " + obj.InterfaceName)
		return

	} else {
		switch obj.Edges.OnLayer.InterfaceLayer {
		case 2:
			if i := usr.QueryOnIPAddress().CountX(context.Background()); i > 0 {
				ipAddr := usr.QueryOnIPAddress().OnlyX(context.Background())
				client.IPAddress.DeleteOne(ipAddr).ExecX(context.Background())
			}

			if obj.Edges.Mode.InterfaceMode == "EtherChannel" {
				response, err = client.NetInterface.UpdateOneID(obj.ID).
					SetInterfaceName(obj.InterfaceName).
					SetInterfaceShutdown(obj.InterfaceShutdown).
					SetOnDevice(obj.Edges.OnDevice).
					SetOnPoInterface(obj.Edges.OnPoInterface).
					ClearOnIPAddress().
					SetOnLayer(obj.Edges.OnLayer).
					SetMode(obj.Edges.Mode).
					Save(context.Background())

			} else if obj.Edges.Mode.InterfaceMode == "None" {

				response, err = client.NetInterface.UpdateOneID(obj.ID).
					SetInterfaceName(obj.InterfaceName).
					SetInterfaceShutdown(obj.InterfaceShutdown).
					SetOnDevice(obj.Edges.OnDevice).
					SetMode(obj.Edges.Mode).
					ClearHaveVlans().
					ClearNativeOnVlan().
					ClearOnIPAddress().
					SetOnLayer(obj.Edges.OnLayer).
					Save(context.Background())

			} else if obj.Edges.NativeOnVlan == nil {
				response, err = client.NetInterface.UpdateOneID(obj.ID).
					SetInterfaceName(obj.InterfaceName).
					SetInterfaceShutdown(obj.InterfaceShutdown).
					SetOnDevice(obj.Edges.OnDevice).
					ClearHaveVlans().
					AddHaveVlans(obj.Edges.HaveVlans...).
					SetMode(obj.Edges.Mode).
					ClearOnIPAddress().
					SetOnLayer(obj.Edges.OnLayer).
					Save(context.Background())

			} else if obj.Edges.HaveVlans == nil {
				response, err = client.NetInterface.UpdateOneID(obj.ID).
					SetInterfaceName(obj.InterfaceName).
					SetInterfaceShutdown(obj.InterfaceShutdown).
					SetOnDevice(obj.Edges.OnDevice).
					SetNativeOnVlan(obj.Edges.NativeOnVlan).
					SetMode(obj.Edges.Mode).
					ClearOnIPAddress().
					SetOnLayer(obj.Edges.OnLayer).
					Save(context.Background())

			} else if obj.Edges.NativeOnVlan == nil && obj.Edges.HaveVlans == nil {
				response, err = client.NetInterface.UpdateOneID(obj.ID).
					SetInterfaceName(obj.InterfaceName).
					SetInterfaceShutdown(obj.InterfaceShutdown).
					SetOnDevice(obj.Edges.OnDevice).
					SetMode(obj.Edges.Mode).
					ClearOnIPAddress().
					SetOnLayer(obj.Edges.OnLayer).
					Save(context.Background())

			} else {
				response, err = client.NetInterface.UpdateOneID(obj.ID).
					SetInterfaceName(obj.InterfaceName).
					SetInterfaceShutdown(obj.InterfaceShutdown).
					SetOnDevice(obj.Edges.OnDevice).
					SetNativeOnVlan(obj.Edges.NativeOnVlan).
					SetMode(obj.Edges.Mode).
					ClearHaveVlans().
					ClearOnIPAddress().
					SetOnLayer(obj.Edges.OnLayer).
					AddHaveVlans(obj.Edges.HaveVlans...).
					Save(context.Background())

			}
		case 3:
			if CheckNetworkOverlap(client, obj.Edges.OnDevice, obj.Edges.OnIPAddress) {
				err = errors.New("network overlap")
				return
			}

			ipAddr, e := CheckNotDuplicateIPAddress(client, obj.Edges.OnDevice, obj.Edges.OnIPAddress)
			if e != nil {

				if i := usr.QueryOnIPAddress().CountX(context.Background()); i > 0 {
					client.IPAddress.Delete().Where(
						ipaddress.And(
							ipaddress.IPAddressEQ(obj.Edges.OnIPAddress.IPAddress),
							ipaddress.SubnetMaskEQ(obj.Edges.OnIPAddress.SubnetMask),
							ipaddress.HasOnDeviceWith(device.IDEQ(obj.Edges.OnDevice.ID)),
						),
					).ExecX(context.Background())
				}

				ipAddr = client.IPAddress.Create().
					SetIPAddress(obj.Edges.OnIPAddress.IPAddress).
					SetSubnetMask(obj.Edges.OnIPAddress.SubnetMask).
					SetOnDevice(obj.Edges.OnDevice).
					SaveX(context.Background())
			} else {
				if ipAddr.ID != obj.Edges.OnIPAddress.ID {
					err = errors.New("ip duplicated")
					return
				}
			}

			modeNone := client.NetInterfaceMode.Query().Where(netinterfacemode.InterfaceModeEQ("None")).OnlyX(context.Background())
			response, err = client.NetInterface.UpdateOneID(obj.ID).
				SetInterfaceName(obj.InterfaceName).
				SetInterfaceShutdown(obj.InterfaceShutdown).
				SetOnDevice(obj.Edges.OnDevice).
				ClearHaveVlans().
				ClearNativeOnVlan().
				ClearOnIPAddress().
				SetMode(modeNone).
				SetOnIPAddress(ipAddr).
				SetOnLayer(obj.Edges.OnLayer).
				Save(context.Background())

		}

		if err != nil {
			response = nil
			return
		}

		if usr.Edges.OnDevice.DeviceCommitConfig {
			_, err = usr.Edges.OnDevice.Update().SetDeviceCommitConfig(false).Save(context.Background())
		}
		return

	}
}

func DeleteInterface(client *ent.Client, id int) (err error) {
	netInt, err := client.NetInterface.Query().Where(netinterface.IDEQ(id)).WithOnDevice().Only(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	netInt.Edges.OnPoInterface = netInt.QueryOnPoInterface().OnlyX(context.Background())
	if netInt.Edges.OnPoInterface != nil {
		err = errors.New("found po interface")
		return
	}

	err = client.NetInterface.DeleteOneID(id).Exec(context.Background())
	if err != nil {
		return
	}

	if netInt.Edges.OnDevice.DeviceCommitConfig {
		_, err := netInt.Edges.OnDevice.Update().SetDeviceCommitConfig(false).Save(context.Background())
		if err != nil {
			log.Fatal(err)
		}
	}
	return
}

func CleanInterface(client *ent.Client, id int) (err error) {
	_, err = client.NetInterface.Delete().Where(netinterface.HasOnDeviceWith(device.IDEQ(id))).Exec(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return
}
