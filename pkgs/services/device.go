package services

import (
	"context"
	"errors"
	"log"

	"github.com/mrzack99s/netcoco/ent"
	"github.com/mrzack99s/netcoco/ent/device"
	"github.com/mrzack99s/netcoco/ent/vlan"
	"github.com/mrzack99s/netcoco/pkgs/security"
)

func GetAllDevice(client *ent.Client) (response []*ent.Device, err error) {
	response, err = client.Device.Query().All(context.Background())
	for _, obj := range response {
		obj.Edges.InType = obj.QueryInType().OnlyX(context.Background())
		obj.Edges.InPlatform = obj.QueryInPlatform().OnlyX(context.Background())
	}
	return
}

func GetDeviceLightWeightQuery(client *ent.Client, id int) (response *ent.Device, err error) {
	response, err = client.Device.Query().Where(device.IDEQ(id)).WithInType().WithInPlatform().WithStoreVlans().
		Only(context.Background())

	return
}

func GetDevicePOIntOnly(client *ent.Client, id int) (response *ent.Device, err error) {
	response, err = client.Device.Query().Where(device.IDEQ(id)).WithInType().WithInPlatform().WithStoreVlans().WithPoInterfaces().
		Only(context.Background())

	if response.Edges.InType.DeviceTypeName == "router" || response.Edges.InType.DeviceTypeName == "l3switch" {
		if i := response.QueryIPStaticRouting().CountX(context.Background()); i > 0 {
			response.Edges.IPStaticRouting = response.QueryIPStaticRouting().AllX(context.Background())
			for _, ipr := range response.Edges.IPStaticRouting {
				ipr.Edges.OnInterface = ipr.QueryOnInterface().OnlyX(context.Background())
			}
		}
	}

	for _, po := range response.Edges.PoInterfaces {
		po.Edges.Mode = po.QueryMode().OnlyX(context.Background())
		po.Edges.OnLayer = po.QueryOnLayer().OnlyX(context.Background())
		if i := po.QueryHaveVlans().CountX(context.Background()); i > 0 {
			po.Edges.HaveVlans = po.QueryHaveVlans().AllX(context.Background())
		}

		if i := po.QueryNativeOnVlan().CountX(context.Background()); i > 0 {
			po.Edges.NativeOnVlan = po.QueryNativeOnVlan().OnlyX(context.Background())
		}

		if po.Edges.OnLayer.InterfaceLayer == 3 {
			po.Edges.OnIPAddress = po.QueryOnIPAddress().OnlyX(context.Background())
		}
	}
	return
}

func GetDevice(client *ent.Client, id int) (response *ent.Device, err error) {
	response, err = client.Device.Query().Where(device.IDEQ(id)).
		WithInterfaces().WithInType().WithInPlatform().WithStoreVlans().WithPoInterfaces().
		Only(context.Background())

	if response.Edges.InType.DeviceTypeName == "router" || response.Edges.InType.DeviceTypeName == "l3switch" {
		if i := response.QueryIPStaticRouting().CountX(context.Background()); i > 0 {
			response.Edges.IPStaticRouting = response.QueryIPStaticRouting().AllX(context.Background())
			for _, ipr := range response.Edges.IPStaticRouting {
				ipr.Edges.OnInterface = ipr.QueryOnInterface().OnlyX(context.Background())
			}
		}
	}

	for _, dd := range response.Edges.Interfaces {
		dd.Edges.Mode = dd.QueryMode().OnlyX(context.Background())
		dd.Edges.OnLayer = dd.QueryOnLayer().OnlyX(context.Background())
		if dd.Edges.Mode.InterfaceMode == "EtherChannel" {
			dd.Edges.OnPoInterface = dd.QueryOnPoInterface().OnlyX(context.Background())
		} else {
			dd.Edges.HaveVlans = dd.QueryHaveVlans().AllX(context.Background())
			if i := dd.QueryNativeOnVlan().CountX(context.Background()); i > 0 {
				dd.Edges.NativeOnVlan = dd.QueryNativeOnVlan().OnlyX(context.Background())
			}
		}
		if dd.Edges.OnLayer.InterfaceLayer == 3 {
			dd.Edges.OnIPAddress = dd.QueryOnIPAddress().OnlyX(context.Background())
		}
	}

	for _, po := range response.Edges.PoInterfaces {
		po.Edges.Mode = po.QueryMode().OnlyX(context.Background())
		po.Edges.OnLayer = po.QueryOnLayer().OnlyX(context.Background())
		if i := po.QueryHaveVlans().CountX(context.Background()); i > 0 {
			po.Edges.HaveVlans = po.QueryHaveVlans().AllX(context.Background())
		}

		if i := po.QueryNativeOnVlan().CountX(context.Background()); i > 0 {
			po.Edges.NativeOnVlan = po.QueryNativeOnVlan().OnlyX(context.Background())
		}

		if po.Edges.OnLayer.InterfaceLayer == 3 {
			po.Edges.OnIPAddress = po.QueryOnIPAddress().OnlyX(context.Background())
		}
	}
	return
}

func CreateDevice(client *ent.Client, obj ent.Device) (response *ent.Device, err error) {
	vlan1, err := client.Vlan.Query().Where(vlan.VlanIDEQ(1)).Only(context.Background())
	if err != nil {
		vlan1, _ = client.Vlan.Create().SetVlanID(1).Save(context.Background())
	}

	if obj.DeviceSSHPort == 0 {
		response, err = client.Device.Create().
			SetDeviceName(obj.DeviceName).
			SetDeviceHostname(obj.DeviceHostname).
			SetDeviceUsername(obj.DeviceUsername).
			SetDevicePassword(security.Encrypt(obj.DevicePassword)).
			SetDeviceSecret(security.Encrypt(obj.DeviceSecret)).
			SetInTypeID(obj.Edges.InType.ID).
			AddStoreVlans(vlan1).
			SetInPlatform(obj.Edges.InPlatform).
			Save(context.Background())
	} else {
		response, err = client.Device.Create().
			SetDeviceName(obj.DeviceName).
			SetDeviceHostname(obj.DeviceHostname).
			SetDeviceUsername(obj.DeviceUsername).
			SetDevicePassword(security.Encrypt(obj.DevicePassword)).
			SetDeviceSSHPort(obj.DeviceSSHPort).
			SetDeviceSecret(security.Encrypt(obj.DeviceSecret)).
			SetInTypeID(obj.Edges.InType.ID).
			SetInPlatform(obj.Edges.InPlatform).
			AddStoreVlans(vlan1).
			Save(context.Background())
	}

	return
}

func AddDeviceVlan(client *ent.Client, obj ent.Device) (response *ent.Device, err error) {
	vlan, err := client.Vlan.Query().Where(vlan.VlanIDEQ(obj.Edges.StoreVlans[0].VlanID)).Only(context.Background())
	if err != nil {
		vlan, _ = client.Vlan.Create().SetVlanID(obj.Edges.StoreVlans[0].VlanID).Save(context.Background())
	}

	response, e := client.Device.UpdateOneID(obj.ID).
		AddStoreVlans(vlan).
		Save(context.Background())

	if e != nil {
		response = nil
		return
	}

	if response.DeviceCommitConfig {
		_, err := response.Update().SetDeviceCommitConfig(false).Save(context.Background())
		if err != nil {
			log.Fatal(err)
		}
	}

	err = nil
	response = &obj
	return

}

func DeleteDeviceVlan(client *ent.Client, id, vid int) error {
	device, err := GetDevice(client, id)
	if err != nil {
		log.Fatal(err)
	}
	device.Edges.Interfaces = device.QueryInterfaces().AllX(context.Background())
	device.Edges.PoInterfaces = device.QueryPoInterfaces().AllX(context.Background())

	foundUsed := false
	for _, d := range device.Edges.Interfaces {
		d.Edges.HaveVlans = d.QueryHaveVlans().AllX(context.Background())
		d.Edges.NativeOnVlan = d.QueryNativeOnVlan().OnlyX(context.Background())

		if d.Edges.NativeOnVlan.VlanID == vid {
			foundUsed = true
		}

		for _, v := range d.Edges.HaveVlans {
			if v.VlanID == vid {
				foundUsed = true
			}
		}
	}

	for _, d := range device.Edges.PoInterfaces {
		d.Edges.HaveVlans = d.QueryHaveVlans().AllX(context.Background())
		d.Edges.NativeOnVlan = d.QueryNativeOnVlan().OnlyX(context.Background())

		if d.Edges.NativeOnVlan.VlanID == vid {
			foundUsed = true
		}

		for _, v := range d.Edges.HaveVlans {
			if v.VlanID == vid {
				foundUsed = true
			}
		}
	}

	if foundUsed {
		return errors.New("vlan used")
	}

	vlan, err := client.Vlan.Query().Where(vlan.VlanIDEQ(vid)).Only(context.Background())
	if err != nil {
		return err
	}

	// Create log
	_, err = client.DeletedVlanLog.Create().SetVlanID(vlan.VlanID).SetOnDevice(device).Save(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.Device.UpdateOneID(id).
		RemoveStoreVlans(vlan).
		Save(context.Background())

	if err != nil {
		return err
	}

	if response.DeviceCommitConfig {
		_, err = response.Update().SetDeviceCommitConfig(false).Save(context.Background())
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil

}

func EditDeviceDetail(client *ent.Client, obj ent.Device) (response *ent.Device, e error) {
	usr, err := client.Device.
		Query().Where(device.IDEQ(obj.ID)).
		Only(context.Background())

	if err != nil {
		response = nil
		e = errors.New("not found device " + obj.DeviceName)
		return

	} else {

		if obj.DevicePassword == "" && obj.DeviceSecret == "" {
			_, e = client.Device.UpdateOneID(usr.ID).
				SetDeviceName(obj.DeviceName).
				SetDeviceHostname(obj.DeviceHostname).
				SetDeviceUsername(obj.DeviceUsername).
				SetDeviceSSHPort(obj.DeviceSSHPort).
				SetInTypeID(obj.Edges.InType.ID).
				SetInPlatform(obj.Edges.InPlatform).
				Save(context.Background())
		} else if obj.DevicePassword == "" {
			newPasswdEncrypt := security.Encrypt(obj.DevicePassword)
			_, e = client.Device.UpdateOneID(usr.ID).
				SetDeviceName(obj.DeviceName).
				SetDeviceHostname(obj.DeviceHostname).
				SetDeviceUsername(obj.DeviceUsername).
				SetDevicePassword(newPasswdEncrypt).
				SetDeviceSSHPort(obj.DeviceSSHPort).
				SetInTypeID(obj.Edges.InType.ID).
				SetInPlatform(obj.Edges.InPlatform).
				Save(context.Background())
		} else if obj.DeviceSecret == "" {
			newSecretEncrypt := security.Encrypt(obj.DeviceSecret)
			_, e = client.Device.UpdateOneID(usr.ID).
				SetDeviceName(obj.DeviceName).
				SetDeviceHostname(obj.DeviceHostname).
				SetDeviceUsername(obj.DeviceUsername).
				SetDeviceSecret(newSecretEncrypt).
				SetDeviceSSHPort(obj.DeviceSSHPort).
				SetInTypeID(obj.Edges.InType.ID).
				SetInPlatform(obj.Edges.InPlatform).
				Save(context.Background())
		} else {
			newSecretEncrypt := security.Encrypt(obj.DeviceSecret)
			newPasswdEncrypt := security.Encrypt(obj.DevicePassword)
			_, e = client.Device.UpdateOneID(usr.ID).
				SetDeviceName(obj.DeviceName).
				SetDeviceHostname(obj.DeviceHostname).
				SetDeviceUsername(obj.DeviceUsername).
				SetDeviceSecret(newSecretEncrypt).
				SetDevicePassword(newPasswdEncrypt).
				SetDeviceSSHPort(obj.DeviceSSHPort).
				SetInTypeID(obj.Edges.InType.ID).
				SetInPlatform(obj.Edges.InPlatform).
				Save(context.Background())
		}

		if e != nil {
			response = nil
			return
		}
		usr.DevicePassword = ""
		usr.DeviceSecret = ""
		response = usr
		return

	}
}

func DeleteDevice(client *ent.Client, id int) (err error) {
	err = client.Device.DeleteOneID(id).Exec(context.Background())
	return
}
