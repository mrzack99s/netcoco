package services

import (
	"context"
	"fmt"

	"github.com/mrzack99s/netcoco/ent"
	"github.com/mrzack99s/netcoco/ent/vlan"
)

func GetAllVlan(client *ent.Client) (response []*ent.Vlan, err error) {
	response, err = client.Vlan.Query().All(context.Background())
	return
}

func GetVlan(client *ent.Client, id int) (response *ent.Vlan, err error) {
	response, err = client.Vlan.Query().Where(vlan.IDEQ(id)).Only(context.Background())
	return
}

func CreateVlan(client *ent.Client, obj ent.Vlan) (response *ent.Vlan, err error) {
	response, err = client.Vlan.Create().
		SetVlanID(obj.VlanID).
		Save(context.Background())

	return
}

func EditVlanDetail(client *ent.Client, obj ent.Vlan) (response *ent.Vlan, e error) {
	usr, err := client.Vlan.
		Query().Where(vlan.IDEQ(obj.ID)).
		Only(context.Background())
	if err != nil {
		response = nil
		e = fmt.Errorf("not found %d", obj.VlanID)
		return

	} else {

		_, e = client.Vlan.UpdateOneID(usr.ID).
			SetVlanID(obj.VlanID).
			Save(context.Background())

		if e != nil {
			response = nil
			return
		}

		response = usr
		return

	}
}

func DeleteVlan(client *ent.Client, id int) (err error) {
	err = client.Vlan.DeleteOneID(id).Exec(context.Background())
	return
}
