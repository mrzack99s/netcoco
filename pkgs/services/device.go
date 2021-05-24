package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/mrzack99s/netcoco/ent"
	"github.com/mrzack99s/netcoco/ent/device"
	"github.com/mrzack99s/netcoco/pkgs/security"
)

func GetAllDevice(client *ent.Client) (response []*ent.Device, err error) {
	response, err = client.Device.Query().All(context.Background())
	for _, obj := range response {
		obj.Edges.InType = obj.QueryInType().OnlyX(context.Background())
	}
	return
}

func GetDevice(client *ent.Client, id int) (response *ent.Device, err error) {
	response, err = client.Device.Query().Where(device.IDEQ(id)).Only(context.Background())
	response.Edges.Interfaces = response.QueryInterfaces().AllX(context.Background())
	for _, dd := range response.Edges.Interfaces {
		dd.Edges.Mode = dd.QueryMode().OnlyX(context.Background())
	}
	response.Edges.InType = response.QueryInType().OnlyX(context.Background())
	return
}

func CreateDevice(client *ent.Client, obj ent.Device) (response *ent.Device, err error) {
	if obj.DeviceSSHPort == 0 {
		response, err = client.Device.Create().
			SetDeviceName(obj.DeviceName).
			SetDeviceHostname(obj.DeviceHostname).
			SetDeviceUsername(obj.DeviceUsername).
			SetDevicePassword(security.Encrypt(obj.DevicePassword)).
			SetDeviceSecret(security.Encrypt(obj.DeviceSecret)).
			SetInTypeID(obj.Edges.InType.ID).
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
			Save(context.Background())
	}

	return
}

func EditDeviceDetail(client *ent.Client, obj ent.Device) (response *ent.Device, e error) {
	usr, err := client.Device.
		Query().Where(device.IDEQ(obj.ID)).
		Only(context.Background())
	if err != nil {
		response = nil
		e = errors.New(fmt.Sprintf("Not found device %s", obj.DeviceName))
		return

	} else {

		if obj.DevicePassword == "" && obj.DeviceSecret == "" {
			_, e = client.Device.UpdateOneID(usr.ID).
				SetDeviceName(obj.DeviceName).
				SetDeviceHostname(obj.DeviceHostname).
				SetDeviceUsername(obj.DeviceUsername).
				SetDeviceSSHPort(obj.DeviceSSHPort).
				SetInTypeID(obj.Edges.InType.ID).
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
	fmt.Println(id)
	err = client.Device.DeleteOneID(id).Exec(context.Background())
	return
}
