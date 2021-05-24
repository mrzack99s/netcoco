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
	if obj.InterfaceNativeVlan == "" {
		response, err = client.NetInterface.Create().
			SetInterfaceName(obj.InterfaceName).
			SetInterfaceVlan(obj.InterfaceVlan).
			SetOnDevice(obj.Edges.OnDevice).
			SetMode(obj.Edges.Mode).
			Save(context.Background())
	} else {
		response, err = client.NetInterface.Create().
			SetInterfaceName(obj.InterfaceName).
			SetInterfaceVlan(obj.InterfaceVlan).
			SetInterfaceNativeVlan(obj.InterfaceNativeVlan).
			SetOnDevice(obj.Edges.OnDevice).
			SetMode(obj.Edges.Mode).
			Save(context.Background())
	}

	return
}

func CreateRangeInterface(client *ent.Client, obj []ent.NetInterface) (response []*ent.NetInterface, err error) {
	for _, item := range obj {
		r, e := client.NetInterface.Create().
			SetInterfaceName(item.InterfaceName).
			SetInterfaceVlan(item.InterfaceVlan).
			SetInterfaceNativeVlan(item.InterfaceNativeVlan).
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

	return
}

func EditInterfaceDetail(client *ent.Client, obj ent.NetInterface) (response *ent.NetInterface, e error) {
	usr, err := client.NetInterface.
		Query().Where(netinterface.IDEQ(obj.ID)).
		Only(context.Background())
	if err != nil {
		response = nil
		e = errors.New(fmt.Sprintf("Not found interface %s", obj.InterfaceName))
		return

	} else {

		if obj.InterfaceNativeVlan == "" {
			_, e = client.NetInterface.UpdateOneID(usr.ID).
				SetInterfaceName(obj.InterfaceName).
				SetInterfaceVlan(obj.InterfaceVlan).
				SetOnDevice(obj.Edges.OnDevice).
				SetMode(obj.Edges.Mode).
				Save(context.Background())
		} else {
			_, e = client.NetInterface.UpdateOneID(usr.ID).
				SetInterfaceName(obj.InterfaceName).
				SetInterfaceVlan(obj.InterfaceVlan).
				SetInterfaceNativeVlan(obj.InterfaceNativeVlan).
				SetOnDevice(obj.Edges.OnDevice).
				SetMode(obj.Edges.Mode).
				Save(context.Background())
		}

		if e != nil {
			response = nil
			return
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
