package services

import (
	"context"
	"errors"
	"log"

	"github.com/mrzack99s/netcoco/ent"
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
	if obj.Edges.NativeOnVlan == nil {
		response, err = client.PortChannelInterface.Create().
			SetPoInterfaceID(obj.PoInterfaceID).
			SetPoInterfaceShutdown(obj.PoInterfaceShutdown).
			AddHaveVlans(obj.Edges.HaveVlans...).
			SetMode(obj.Edges.Mode).
			SetOnDevice(obj.Edges.OnDevice).
			Save(context.Background())

	} else if obj.Edges.HaveVlans == nil {
		response, err = client.PortChannelInterface.Create().
			SetPoInterfaceID(obj.PoInterfaceID).
			SetPoInterfaceShutdown(obj.PoInterfaceShutdown).
			SetNativeOnVlan(obj.Edges.NativeOnVlan).
			SetMode(obj.Edges.Mode).
			SetOnDevice(obj.Edges.OnDevice).
			Save(context.Background())

	} else if obj.Edges.NativeOnVlan == nil && obj.Edges.HaveVlans == nil {
		response, err = client.PortChannelInterface.Create().
			SetPoInterfaceID(obj.PoInterfaceID).
			SetPoInterfaceShutdown(obj.PoInterfaceShutdown).
			SetMode(obj.Edges.Mode).
			SetOnDevice(obj.Edges.OnDevice).
			Save(context.Background())

	} else {
		response, err = client.PortChannelInterface.Create().
			SetPoInterfaceID(obj.PoInterfaceID).
			SetPoInterfaceShutdown(obj.PoInterfaceShutdown).
			SetNativeOnVlan(obj.Edges.NativeOnVlan).
			SetMode(obj.Edges.Mode).
			SetOnDevice(obj.Edges.OnDevice).
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

func EditPortChannelInterface(client *ent.Client, obj ent.PortChannelInterface) (response *ent.PortChannelInterface, err error) {
	usr, err := client.PortChannelInterface.
		Query().Where(portchannelinterface.IDEQ(obj.ID)).
		Only(context.Background())
	if err != nil {
		response = nil
		err = errors.New("not found interface ")
		return

	} else {
		if obj.Edges.NativeOnVlan == nil {
			response, err = client.PortChannelInterface.UpdateOneID(usr.ID).
				SetPoInterfaceID(obj.PoInterfaceID).
				SetPoInterfaceShutdown(obj.PoInterfaceShutdown).
				ClearHaveVlans().
				SetOnDevice(obj.Edges.OnDevice).
				AddHaveVlans(obj.Edges.HaveVlans...).
				SetMode(obj.Edges.Mode).
				Save(context.Background())

		} else if obj.Edges.HaveVlans == nil {
			response, err = client.PortChannelInterface.UpdateOneID(usr.ID).
				SetPoInterfaceID(obj.PoInterfaceID).
				SetPoInterfaceShutdown(obj.PoInterfaceShutdown).
				SetNativeOnVlan(obj.Edges.NativeOnVlan).
				SetMode(obj.Edges.Mode).
				SetOnDevice(obj.Edges.OnDevice).
				Save(context.Background())

		} else if obj.Edges.NativeOnVlan == nil && obj.Edges.HaveVlans == nil {
			response, err = client.PortChannelInterface.UpdateOneID(usr.ID).
				SetPoInterfaceID(obj.PoInterfaceID).
				SetPoInterfaceShutdown(obj.PoInterfaceShutdown).
				SetMode(obj.Edges.Mode).
				SetOnDevice(obj.Edges.OnDevice).
				Save(context.Background())

		} else {
			response, err = client.PortChannelInterface.UpdateOneID(usr.ID).
				SetPoInterfaceID(obj.PoInterfaceID).
				SetPoInterfaceShutdown(obj.PoInterfaceShutdown).
				SetNativeOnVlan(obj.Edges.NativeOnVlan).
				SetMode(obj.Edges.Mode).
				ClearHaveVlans().
				AddHaveVlans(obj.Edges.HaveVlans...).
				SetOnDevice(obj.Edges.OnDevice).
				Save(context.Background())

		}

		if err != nil {
			response = nil
			return
		}

		if obj.Edges.OnDevice.DeviceCommitConfig {
			_, err = obj.Edges.OnDevice.Update().SetDeviceCommitConfig(false).Save(context.Background())
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
			SetNativeOnVlan(po.Edges.NativeOnVlan).
			AddHaveVlans(po.Edges.HaveVlans...).
			SetInterfaceShutdown(true).
			Save(context.Background())
		if err != nil {
			log.Fatal(err)
		}
	}

	err = client.PortChannelInterface.DeleteOneID(id).Exec(context.Background())
	return
}
