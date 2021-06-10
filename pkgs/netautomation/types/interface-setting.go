package types

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mrzack99s/netcoco/ent"
)

type InterfaceL2Setting struct {
	InterfaceName string
	PortChannelID int
	Mode          string
	VLANs         string
	NativeVLAN    string
	Shutdown      bool
	PoSetting     bool
}

type InterfaceL3Setting struct {
	InterfaceName string
	DeviceType    string
	Shutdown      bool
	IPAddress     string
	SubnetMask    string
}

func GetHaveVlans(vlans []*ent.Vlan) (intVlan strings.Builder) {
	intVlan = strings.Builder{}
	if len(vlans) > 0 {
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
