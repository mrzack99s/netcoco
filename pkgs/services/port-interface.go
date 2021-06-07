package services

import (
	"context"
	"errors"
	"log"

	"github.com/mrzack99s/netcoco/ent"
	"github.com/mrzack99s/netcoco/ent/device"
	"github.com/mrzack99s/netcoco/ent/ipaddress"
	"github.com/mrzack99s/netcoco/ent/netinterfacemode"
	"github.com/mrzack99s/netcoco/ent/portchannelinterface"
)

func GetPortChannelInterface(client *ent.Client, id int) (response *ent.PortChannelInterface, err error) {
	response, err = client.PortChannelInterface.Query().Where(portchannelinterface.IDEQ(id)).Only(context.Background())
	response.Edges.Mode = response.QueryMode().OnlyX(context.Background())
	response.Edges.HaveVlans = response.QueryHaveVlans().AllX(context.Background())
	response.Edges.NativeOnVlan = response.QueryHaveVlans().OnlyX(context.Background())
	return
}

func CreatePortChannelInterface(client *ent.Client, obj ent.PortChannelInterface) (response *ent.PortChannelInterface, err error) {
	switch obj.Edges.OnLayer.InterfaceLayer {
	case 2:
		if obj.Edges.NativeOnVlan == nil {
			response, err = client.PortChannelInterface.Create().
				SetPoInterfaceID(obj.PoInterfaceID).
				SetPoInterfaceShutdown(obj.PoInterfaceShutdown).
				AddHaveVlans(obj.Edges.HaveVlans...).
				SetMode(obj.Edges.Mode).
				SetOnDevice(obj.Edges.OnDevice).
				SetOnLayer(obj.Edges.OnLayer).
				Save(context.Background())

		} else if obj.Edges.HaveVlans == nil {
			response, err = client.PortChannelInterface.Create().
				SetPoInterfaceID(obj.PoInterfaceID).
				SetPoInterfaceShutdown(obj.PoInterfaceShutdown).
				SetNativeOnVlan(obj.Edges.NativeOnVlan).
				SetMode(obj.Edges.Mode).
				SetOnDevice(obj.Edges.OnDevice).
				SetOnLayer(obj.Edges.OnLayer).
				Save(context.Background())

		} else if obj.Edges.NativeOnVlan == nil && obj.Edges.HaveVlans == nil {
			response, err = client.PortChannelInterface.Create().
				SetPoInterfaceID(obj.PoInterfaceID).
				SetPoInterfaceShutdown(obj.PoInterfaceShutdown).
				SetMode(obj.Edges.Mode).
				SetOnDevice(obj.Edges.OnDevice).
				SetOnLayer(obj.Edges.OnLayer).
				Save(context.Background())

		} else {
			response, err = client.PortChannelInterface.Create().
				SetPoInterfaceID(obj.PoInterfaceID).
				SetPoInterfaceShutdown(obj.PoInterfaceShutdown).
				SetNativeOnVlan(obj.Edges.NativeOnVlan).
				SetMode(obj.Edges.Mode).
				SetOnLayer(obj.Edges.OnLayer).
				SetOnDevice(obj.Edges.OnDevice).
				AddHaveVlans(obj.Edges.HaveVlans...).
				Save(context.Background())

		}
	case 3:
		if !CheckNetworkOverlap(client, obj.Edges.OnDevice, obj.Edges.OnIPAddress) {
			err = errors.New("network overlap")
			return
		}

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
		response, err = client.PortChannelInterface.Create().
			SetPoInterfaceID(obj.PoInterfaceID).
			SetPoInterfaceShutdown(obj.PoInterfaceShutdown).
			SetOnDevice(obj.Edges.OnDevice).
			SetOnIPAddress(ipAddr).
			SetOnLayer(obj.Edges.OnLayer).
			SetMode(modeNone).
			Save(context.Background())

	}

	response.Edges.OnDevice = response.QueryOnDevice().OnlyX(context.Background())
	if response.Edges.OnDevice.DeviceCommitConfig {
		_, err = response.Edges.OnDevice.Update().SetDeviceCommitConfig(false).Save(context.Background())
		if err != nil {
			log.Fatal(err)
		}
	}

	return
}

func EditPortChannelInterface(client *ent.Client, obj ent.PortChannelInterface) (response *ent.PortChannelInterface, err error) {
	usr, err := client.PortChannelInterface.
		Query().Where(portchannelinterface.IDEQ(obj.ID)).
		Only(context.Background())
	if err != nil {
		response = nil
		err = errors.New("not found interface ")
		return

	} else {
		switch obj.Edges.OnLayer.InterfaceLayer {
		case 2:
			if i := usr.QueryOnIPAddress().CountX(context.Background()); i > 0 {
				ipAddr := usr.QueryOnIPAddress().OnlyX(context.Background())
				client.IPAddress.DeleteOne(ipAddr).ExecX(context.Background())
			}

			if obj.Edges.Mode.InterfaceMode == "None" {
				response, err = client.PortChannelInterface.UpdateOneID(usr.ID).
					SetPoInterfaceID(obj.PoInterfaceID).
					SetPoInterfaceShutdown(obj.PoInterfaceShutdown).
					ClearHaveVlans().
					ClearNativeOnVlan().
					SetOnDevice(obj.Edges.OnDevice).
					SetMode(obj.Edges.Mode).
					ClearOnIPAddress().
					SetOnLayer(obj.Edges.OnLayer).
					Save(context.Background())

			} else if obj.Edges.NativeOnVlan == nil {
				response, err = client.PortChannelInterface.UpdateOneID(usr.ID).
					SetPoInterfaceID(obj.PoInterfaceID).
					SetPoInterfaceShutdown(obj.PoInterfaceShutdown).
					ClearHaveVlans().
					SetOnDevice(obj.Edges.OnDevice).
					AddHaveVlans(obj.Edges.HaveVlans...).
					SetMode(obj.Edges.Mode).
					ClearOnIPAddress().
					SetOnLayer(obj.Edges.OnLayer).
					Save(context.Background())

			} else if obj.Edges.HaveVlans == nil {
				response, err = client.PortChannelInterface.UpdateOneID(usr.ID).
					SetPoInterfaceID(obj.PoInterfaceID).
					SetPoInterfaceShutdown(obj.PoInterfaceShutdown).
					SetNativeOnVlan(obj.Edges.NativeOnVlan).
					SetMode(obj.Edges.Mode).
					SetOnDevice(obj.Edges.OnDevice).
					ClearOnIPAddress().
					SetOnLayer(obj.Edges.OnLayer).
					Save(context.Background())

			} else if obj.Edges.NativeOnVlan == nil && obj.Edges.HaveVlans == nil {
				response, err = client.PortChannelInterface.UpdateOneID(usr.ID).
					SetPoInterfaceID(obj.PoInterfaceID).
					SetPoInterfaceShutdown(obj.PoInterfaceShutdown).
					SetMode(obj.Edges.Mode).
					SetOnDevice(obj.Edges.OnDevice).
					ClearOnIPAddress().
					SetOnLayer(obj.Edges.OnLayer).
					Save(context.Background())

			} else {
				response, err = client.PortChannelInterface.UpdateOneID(usr.ID).
					SetPoInterfaceID(obj.PoInterfaceID).
					SetPoInterfaceShutdown(obj.PoInterfaceShutdown).
					SetNativeOnVlan(obj.Edges.NativeOnVlan).
					SetMode(obj.Edges.Mode).
					ClearHaveVlans().
					ClearOnIPAddress().
					SetOnLayer(obj.Edges.OnLayer).
					AddHaveVlans(obj.Edges.HaveVlans...).
					SetOnDevice(obj.Edges.OnDevice).
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

			response, err = client.PortChannelInterface.UpdateOneID(obj.ID).
				SetPoInterfaceID(obj.PoInterfaceID).
				SetPoInterfaceShutdown(obj.PoInterfaceShutdown).
				SetOnDevice(obj.Edges.OnDevice).
				ClearHaveVlans().
				ClearNativeOnVlan().
				ClearOnIPAddress().
				SetOnIPAddress(ipAddr).
				SetOnLayer(obj.Edges.OnLayer).
				Save(context.Background())

		}

		if err != nil {
			response = nil
			return
		}

		response.Edges.OnDevice = response.QueryOnDevice().OnlyX(context.Background())
		if response.Edges.OnDevice.DeviceCommitConfig {
			_, err = response.Edges.OnDevice.Update().SetDeviceCommitConfig(false).Save(context.Background())
			if err != nil {
				log.Fatal(err)
			}
		}

	}
	return
}

func DeletePortChannelInterface(client *ent.Client, id int) (err error) {
	po, err := client.PortChannelInterface.Query().Where(portchannelinterface.IDEQ(id)).Only(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	modeNone, err := client.NetInterfaceMode.Query().Where(netinterfacemode.InterfaceModeEQ("None")).Only(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	po.Edges.Interfaces = po.QueryInterfaces().AllX(context.Background())
	po.Edges.HaveVlans = po.QueryHaveVlans().AllX(context.Background())
	po.Edges.NativeOnVlan = po.QueryNativeOnVlan().OnlyX(context.Background())
	for _, i := range po.Edges.Interfaces {
		_, err = i.Update().
			SetMode(modeNone).
			ClearHaveVlans().
			ClearNativeOnVlan().
			SetInterfaceShutdown(true).
			Save(context.Background())
		if err != nil {
			log.Fatal(err)
		}
	}
	po.Edges.OnDevice = po.QueryOnDevice().OnlyX(context.Background())
	if po.Edges.OnDevice.DeviceCommitConfig {
		_, err = po.Edges.OnDevice.Update().SetDeviceCommitConfig(false).Save(context.Background())
		if err != nil {
			log.Fatal(err)
		}
	}

	err = client.PortChannelInterface.DeleteOneID(id).Exec(context.Background())

	return
}
