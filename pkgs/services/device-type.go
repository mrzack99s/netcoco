package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/mrzack99s/netcoco/ent"
	"github.com/mrzack99s/netcoco/ent/devicetype"
)

func GetAllDeviceType(client *ent.Client) (response []*ent.DeviceType, err error) {
	response, err = client.DeviceType.Query().All(context.Background())
	return
}

func GetDeviceType(client *ent.Client, id int) (response *ent.DeviceType, err error) {
	response, err = client.DeviceType.Query().Where(devicetype.IDEQ(id)).Only(context.Background())
	return
}

func CreateDeviceType(client *ent.Client, obj ent.DeviceType) (response *ent.DeviceType, err error) {
	response, err = client.DeviceType.Create().
		SetDeviceTypeName(obj.DeviceTypeName).
		Save(context.Background())

	return
}

func EditDeviceTypeDetail(client *ent.Client, obj ent.DeviceType) (response *ent.DeviceType, e error) {
	usr, err := client.DeviceType.
		Query().Where(devicetype.IDEQ(obj.ID)).
		Only(context.Background())
	if err != nil {
		response = nil
		e = errors.New(fmt.Sprintf("Not found device %s", obj.DeviceTypeName))
		return

	} else {

		_, e = client.DeviceType.UpdateOneID(usr.ID).
			SetDeviceTypeName(obj.DeviceTypeName).
			Save(context.Background())

		if e != nil {
			response = nil
			return
		}

		response = usr
		return

	}
}

func DeleteDeviceType(client *ent.Client, id int) (err error) {
	err = client.DeviceType.DeleteOneID(id).Exec(context.Background())
	return
}
