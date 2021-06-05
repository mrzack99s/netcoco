package types

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mrzack99s/netcoco/ent"
)

type InterfaceSetting struct {
	InterfaceName string
	PortChannelID int
	Mode          string
	VLANs         string
	NativeVLAN    string
	Shutdown      bool
	PoSetting     bool
}

func GetHaveVlans(vlans []*ent.Vlan) (intVlan strings.Builder) {
	intVlan = strings.Builder{}
	if vlans != nil {
		if len(vlans) > 1 {
			for index, v := range vlans {
				if index == len(vlans)-1 {
					intVlan.WriteString(fmt.Sprintf("%d", v.VlanID))
				} else {
					intVlan.WriteString(fmt.Sprintf("%d,", v.VlanID))
				}

			}
		} else {
			intVlan.WriteString(fmt.Sprintf("%d", vlans[0].VlanID))
		}
	}
	return
}

func GetNativeVlan(nativeVlan *ent.Vlan) (nVlan string) {
	if nativeVlan != nil {
		nVlan = strconv.Itoa(nativeVlan.VlanID)
	}

	return
}
