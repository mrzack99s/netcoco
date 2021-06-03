package sg350

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/mrzack99s/netcoco/ent"
)

func getVlanConfig(vlans string) (config []string) {

	var rawCmd string = `vlan %s`

	replaceRawCmd := fmt.Sprintf(rawCmd, vlans)
	config = strings.Split(replaceRawCmd, "\n")

	return
}

func getDeletesVlanConfig(vlans string) (config []string) {

	var rawCmd string = `no vlan %s`

	replaceRawCmd := fmt.Sprintf(rawCmd, vlans)
	config = strings.Split(replaceRawCmd, "\n")

	return
}

func GetCommitVlanConfig(device *ent.Device) (config []string) {

	device.Edges.StoreVlans = device.QueryStoreVlans().AllX(context.Background())
	device.Edges.DeletedVlans = device.QueryDeletedVlans().AllX(context.Background())

	vlans := strings.Builder{}
	if len(device.Edges.StoreVlans) > 1 {
		for index, o := range device.Edges.StoreVlans {
			if index == len(device.Edges.StoreVlans)-1 {
				vlans.WriteString(fmt.Sprintf("%d", o.VlanID))
			} else {
				vlans.WriteString(fmt.Sprintf("%d,", o.VlanID))
			}
		}
	} else {
		vlans.WriteString(fmt.Sprintf("%d", device.Edges.StoreVlans[0].VlanID))
	}

	config = append(config, getVlanConfig(vlans.String())...)

	dvlans := strings.Builder{}
	if len(device.Edges.DeletedVlans) > 0 {
		if len(device.Edges.DeletedVlans) > 1 {
			for index, o := range device.Edges.DeletedVlans {
				if index == len(device.Edges.DeletedVlans)-1 {
					dvlans.WriteString(fmt.Sprintf("%d", o.VlanID))
				} else {
					dvlans.WriteString(fmt.Sprintf("%d,", o.VlanID))
				}
			}
		} else {
			dvlans.WriteString(fmt.Sprintf("%d", device.Edges.DeletedVlans[0].VlanID))
		}

		config = append(config, getDeletesVlanConfig(dvlans.String())...)

		allDeleteVlan := device.QueryDeletedVlans().AllX(context.Background())
		for _, dd := range allDeleteVlan {
			_, err := dd.Update().SetDeleted(true).Save(context.Background())
			if err != nil {
				log.Fatal(err)
			}
		}
		_, err := device.Update().ClearDeletedVlans().Save(context.Background())
		if err != nil {
			log.Fatal(err)
		}
	}

	return
}
