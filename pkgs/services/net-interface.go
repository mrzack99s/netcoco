package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/mrzack99s/netcoco/ent"
	"github.com/mrzack99s/netcoco/ent/device"
	"github.com/mrzack99s/netcoco/ent/netinterface"
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
		obj.Edges.OnDevice.Update().SetDeviceCommitConfig(false).Save(context.Background())
	}

	return
}

func CreateRangeInterface(client *ent.Client, obj []ent.NetInterface) (response []*ent.NetInterface, err error) {
	for _, item := range obj {
		r, e := client.NetInterface.Create().
			SetInterfaceName(item.InterfaceName).
			SetOnDevice(item.Edges.OnDevice).
			SetMode(item.Edges.Mode).
			Save(context.Background())
		if e != nil {
			err = e
			return
		} else {
			response = append(response, r)
		}
	}
	obj[0].Edges.OnDevice = obj[0].QueryOnDevice().OnlyX(context.Background())
	if obj[0].Edges.OnDevice.DeviceCommitConfig {
		obj[0].Edges.OnDevice.Update().SetDeviceCommitConfig(false).Save(context.Background())
	}

	return
}

func EditInterfaceDetail(client *ent.Client, obj ent.NetInterface) (response *ent.NetInterface, e error) {
	usr, err := client.NetInterface.
		Query().Where(netinterface.IDEQ(obj.ID)).
		Only(context.Background())
	usr.Edges.OnDevice = usr.QueryOnDevice().OnlyX(context.Background())
	if err != nil {
		response = nil
		e = errors.New(fmt.Sprintf("Not found interface %s", obj.InterfaceName))
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

		if e != nil {
			response = nil
			return
		}

		if usr.Edges.OnDevice.DeviceCommitConfig {
			usr.Edges.OnDevice.Update().SetDeviceCommitConfig(false).Save(context.Background())
		}
		response = usr
		return

	}
}

func DeleteInterface(client *ent.Client, id int) (err error) {
	err = client.NetInterface.DeleteOneID(id).Exec(context.Background())
	return
}

func CleanInterface(client *ent.Client, id int) (err error) {
	_, err = client.NetInterface.Delete().Where(netinterface.HasOnDeviceWith(device.IDEQ(id))).Exec(context.Background())
	return
}
