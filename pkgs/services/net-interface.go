package services

import (
	"context"
	"errors"
	"log"

	"github.com/mrzack99s/netcoco/ent"
	"github.com/mrzack99s/netcoco/ent/device"
	"github.com/mrzack99s/netcoco/ent/netinterface"
	"github.com/mrzack99s/netcoco/ent/vlan"
)

func GetInterface(client *ent.Client, id int) (response *ent.NetInterface, err error) {
	response, err = client.NetInterface.Query().Where(netinterface.IDEQ(id)).Only(context.Background())
	response.Edges.Mode = response.QueryMode().OnlyX(context.Background())
	return
}

func CreateInterface(client *ent.Client, obj ent.NetInterface) (response *ent.NetInterface, err error) {
	if obj.Edges.NativeOnVlan == nil {
		response, err = client.NetInterface.Create().
			SetInterfaceName(obj.InterfaceName).
			SetInterfaceShutdown(obj.InterfaceShutdown).
			SetOnDevice(obj.Edges.OnDevice).
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
			Save(context.Background())

	} else if obj.Edges.NativeOnVlan == nil && obj.Edges.HaveVlans == nil {
		response, err = client.NetInterface.Create().
			SetInterfaceName(obj.InterfaceName).
			SetInterfaceShutdown(obj.InterfaceShutdown).
			SetOnDevice(obj.Edges.OnDevice).
			SetMode(obj.Edges.Mode).
			Save(context.Background())

	} else {
		response, err = client.NetInterface.Create().
			SetInterfaceName(obj.InterfaceName).
			SetInterfaceShutdown(obj.InterfaceShutdown).
			SetOnDevice(obj.Edges.OnDevice).
			SetNativeOnVlan(obj.Edges.NativeOnVlan).
			SetMode(obj.Edges.Mode).
			AddHaveVlans(obj.Edges.HaveVlans...).
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

	for _, item := range obj {
		r, e := client.NetInterface.Create().
			SetInterfaceName(item.InterfaceName).
			SetOnDevice(item.Edges.OnDevice).
			SetNativeOnVlan(vlan1).
			SetMode(item.Edges.Mode).
			AddHaveVlans(vlan1).
			Save(context.Background())
		if e != nil {
			err = e
			return
		} else {
			response = append(response, r)
		}
	}

	response[0].Edges.OnDevice = response[0].QueryOnDevice().OnlyX(context.Background())
	if response[0].Edges.OnDevice.DeviceCommitConfig {
		_, err = response[0].Edges.OnDevice.Update().SetDeviceCommitConfig(false).Save(context.Background())
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

		if obj.Edges.NativeOnVlan == nil {
			response, err = client.NetInterface.UpdateOneID(obj.ID).
				SetInterfaceName(obj.InterfaceName).
				SetInterfaceShutdown(obj.InterfaceShutdown).
				SetOnDevice(obj.Edges.OnDevice).
				ClearHaveVlans().
				AddHaveVlans(obj.Edges.HaveVlans...).
				SetMode(obj.Edges.Mode).
				Save(context.Background())

		} else if obj.Edges.HaveVlans == nil {
			response, err = client.NetInterface.UpdateOneID(obj.ID).
				SetInterfaceName(obj.InterfaceName).
				SetInterfaceShutdown(obj.InterfaceShutdown).
				SetOnDevice(obj.Edges.OnDevice).
				SetNativeOnVlan(obj.Edges.NativeOnVlan).
				SetMode(obj.Edges.Mode).
				Save(context.Background())

		} else if obj.Edges.NativeOnVlan == nil && obj.Edges.HaveVlans == nil {
			response, err = client.NetInterface.UpdateOneID(obj.ID).
				SetInterfaceName(obj.InterfaceName).
				SetInterfaceShutdown(obj.InterfaceShutdown).
				SetOnDevice(obj.Edges.OnDevice).
				SetMode(obj.Edges.Mode).
				Save(context.Background())

		} else {
			response, err = client.NetInterface.UpdateOneID(obj.ID).
				SetInterfaceName(obj.InterfaceName).
				SetInterfaceShutdown(obj.InterfaceShutdown).
				SetOnDevice(obj.Edges.OnDevice).
				SetNativeOnVlan(obj.Edges.NativeOnVlan).
				SetMode(obj.Edges.Mode).
				ClearHaveVlans().
				AddHaveVlans(obj.Edges.HaveVlans...).
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
	err = client.NetInterface.DeleteOneID(id).Exec(context.Background())
	return
}

func CleanInterface(client *ent.Client, id int) (err error) {
	_, err = client.NetInterface.Delete().Where(netinterface.HasOnDeviceWith(device.IDEQ(id))).Exec(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return
}
