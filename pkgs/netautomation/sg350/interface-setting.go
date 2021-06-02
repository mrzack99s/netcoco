package sg350

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/mrzack99s/netcoco/ent"
	"github.com/mrzack99s/netcoco/pkgs/netautomation/constants"
	"github.com/mrzack99s/netcoco/pkgs/netautomation/types"
)

func getInterfaceConfig(o *types.InterfaceSetting) (config []string) {

	switch o.Shutdown {
	case true:
		switch o.Mode {
		case constants.MODE_ACCESS:
			var rawCmd string = `interface %s
switchport mode access
no switchport trunk native vlan
no switchport trunk allowed vlan
switchport access vlan %s
shutdown
exit`
			replaceRawCmd := fmt.Sprintf(rawCmd, o.InterfaceName, o.VLANs)
			config = strings.Split(replaceRawCmd, "\n")
		case constants.MODE_TRUNKING:
			var rawCmd string = `interface %s
switchport mode trunk
no switchport access vlan
switchport trunk allowed vlan %s
switchport trunk native vlan %s
shutdown
exit`
			if o.NativeVLAN == "1" {
				rawCmd = `interface %s
switchport mode trunk
no switchport access vlan
no switchport trunk native vlan
switchport trunk allowed vlan add %s
shutdown
exit`
				replaceRawCmd := fmt.Sprintf(rawCmd, o.InterfaceName, o.VLANs)
				config = strings.Split(replaceRawCmd, "\n")
			} else {
				replaceRawCmd := fmt.Sprintf(rawCmd, o.InterfaceName, o.VLANs, o.NativeVLAN)
				config = strings.Split(replaceRawCmd, "\n")
			}

		case constants.MODE_NONE:
			var rawCmd string = `interface %s
no switchport mode
no switchport access vlan
no switchport trunk native vlan
no sw trunk allowed vlan
shutdown
exit`
			replaceRawCmd := fmt.Sprintf(rawCmd, o.InterfaceName)
			config = strings.Split(replaceRawCmd, "\n")

		}
	case false:
		switch o.Mode {
		case constants.MODE_ACCESS:
			var rawCmd string = `interface %s
switchport mode access
no switchport trunk native vlan
no switchport trunk allowed vlan
switchport access vlan %s
no shutdown
exit`
			replaceRawCmd := fmt.Sprintf(rawCmd, o.InterfaceName, o.VLANs)
			config = strings.Split(replaceRawCmd, "\n")
		case constants.MODE_TRUNKING:
			var rawCmd string = `interface %s
switchport mode trunk
no switchport access vlan
switchport trunk allowed vlan %s
switchport trunk native vlan %s
no shutdown
exit`
			if o.NativeVLAN == "1" {
				rawCmd = `interface %s
switchport mode trunk
no switchport access vlan
switchport trunk allowed vlan %s
no switchport trunk native vlan
no shutdown
exit`
				replaceRawCmd := fmt.Sprintf(rawCmd, o.InterfaceName, o.VLANs)
				config = strings.Split(replaceRawCmd, "\n")
			} else {
				replaceRawCmd := fmt.Sprintf(rawCmd, o.InterfaceName, o.VLANs, o.NativeVLAN)
				config = strings.Split(replaceRawCmd, "\n")
			}

		case constants.MODE_NONE:
			var rawCmd string = `interface %s
no switchport mode
no switchport access vlan
no switchport trunk native vlan
no sw trunk allowed vlan
no shutdown
exit`
			replaceRawCmd := fmt.Sprintf(rawCmd, o.InterfaceName)
			config = strings.Split(replaceRawCmd, "\n")

		}
	}

	return
}

func GetCommitInterfaceConfig(device *ent.Device) (config []string) {
	for _, obj := range device.Edges.Interfaces {

		obj.Edges.Mode = obj.QueryMode().OnlyX(context.Background())
		obj.Edges.HaveVlans = obj.QueryHaveVlans().AllX(context.Background())
		obj.Edges.NativeOnVlan = obj.QueryNativeOnVlan().OnlyX(context.Background())

		mode := constants.MODE_NONE
		if obj.Edges.Mode.InterfaceMode == "Access" {
			mode = constants.MODE_ACCESS
		} else if obj.Edges.Mode.InterfaceMode == "Trunking" {
			mode = constants.MODE_TRUNKING
		}

		intVlan := strings.Builder{}
		if obj.Edges.HaveVlans != nil {
			if len(obj.Edges.HaveVlans) > 1 {
				for index, v := range obj.Edges.HaveVlans {
					if index == len(obj.Edges.HaveVlans)-1 {
						intVlan.WriteString(fmt.Sprintf("%d", v.VlanID))
					} else {
						intVlan.WriteString(fmt.Sprintf("%d,", v.VlanID))
					}

				}
			} else {
				intVlan.WriteString(fmt.Sprintf("%d", obj.Edges.HaveVlans[0].VlanID))
			}
		}

		nVlan := ""
		if obj.Edges.NativeOnVlan != nil {
			nVlan = strconv.FormatInt(int64(obj.Edges.NativeOnVlan.VlanID), 10)
		}

		intSetting := &types.InterfaceSetting{
			InterfaceName: obj.InterfaceName,
			Mode:          mode,
			VLANs:         intVlan.String(),
			NativeVLAN:    nVlan,
			Shutdown:      obj.InterfaceShutdown,
		}

		config = append(config, getInterfaceConfig(intSetting)...)
	}
	return
}
