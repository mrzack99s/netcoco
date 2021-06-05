package ios

import (
	"context"
	"fmt"
	"strings"

	"github.com/mrzack99s/netcoco/ent"
	"github.com/mrzack99s/netcoco/pkgs/netautomation/constants"
	"github.com/mrzack99s/netcoco/pkgs/netautomation/types"
)

func getInterfaceConfig(o *types.InterfaceSetting) (config []string) {

	switch o.Shutdown {
	case true:
		switch o.Mode {
		case constants.MODE_ETHERCHANNEL:
			var rawCmd string = `interface %s
channel-group %d mode active
shutdown
exit`
			replaceRawCmd := fmt.Sprintf(rawCmd, o.InterfaceName, o.PortChannelID)
			config = strings.Split(replaceRawCmd, "\n")
		case constants.MODE_ACCESS:
			var rawCmd string = `interface %s
switchport mode access
no switchport trunk allowed vlan
no switchport trunk native vlan
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
switchport trunk allowed vlan %s
no switchport trunk native vlan
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
no switchport trunk allowed vlan
no switchport trunk native vlan
shutdown
exit`
			replaceRawCmd := fmt.Sprintf(rawCmd, o.InterfaceName)
			config = strings.Split(replaceRawCmd, "\n")

		}
	case false:
		switch o.Mode {
		case constants.MODE_ETHERCHANNEL:
			var rawCmd string = `interface %s
channel-group %d mode active
no shutdown
exit`
			replaceRawCmd := fmt.Sprintf(rawCmd, o.InterfaceName, o.PortChannelID)
			config = strings.Split(replaceRawCmd, "\n")
		case constants.MODE_ACCESS:
			var rawCmd string = `interface %s
switchport mode access
no switchport trunk allowed vlan
no switchport trunk native vlan
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
no switchport trunk allowed vlan
no switchport trunk native vlan
no shutdown
exit`
			replaceRawCmd := fmt.Sprintf(rawCmd, o.InterfaceName)
			config = strings.Split(replaceRawCmd, "\n")

		}
	}

	if !o.PoSetting {
		tmp0 := config[0]
		tmp1 := config[1:]
		config = nil
		config = append(config, tmp0)
		config = append(config, "no channel-group")
		config = append(config, tmp1...)
	}

	return
}

func GetCommitInterfaceConfig(device *ent.Device) (config []string) {
	var portChannelConfig []string

	for _, obj := range device.Edges.Interfaces {

		obj.Edges.Mode = obj.QueryMode().OnlyX(context.Background())

		mode := constants.GetIntMode(obj.Edges.Mode)

		if mode == constants.MODE_ETHERCHANNEL {
			obj.Edges.OnPoInterface = obj.QueryOnPoInterface().OnlyX(context.Background())
			obj.Edges.OnPoInterface.Edges.Mode = obj.Edges.OnPoInterface.QueryMode().OnlyX(context.Background())
			obj.Edges.OnPoInterface.Edges.HaveVlans = obj.Edges.OnPoInterface.QueryHaveVlans().AllX(context.Background())
			obj.Edges.OnPoInterface.Edges.NativeOnVlan = obj.Edges.OnPoInterface.QueryNativeOnVlan().OnlyX(context.Background())
		} else {
			obj.Edges.HaveVlans = obj.QueryHaveVlans().AllX(context.Background())
			obj.Edges.NativeOnVlan = obj.QueryNativeOnVlan().OnlyX(context.Background())
		}

		if mode == constants.MODE_ETHERCHANNEL {
			intSetting := &types.InterfaceSetting{
				InterfaceName: obj.InterfaceName,
				PortChannelID: obj.Edges.OnPoInterface.PoInterfaceID,
				Mode:          mode,
				PoSetting:     true,
				Shutdown:      obj.InterfaceShutdown,
			}
			config = append(config, getInterfaceConfig(intSetting)...)

			intVlan := types.GetHaveVlans(obj.Edges.OnPoInterface.Edges.HaveVlans)
			nVlan := types.GetNativeVlan(obj.Edges.OnPoInterface.Edges.NativeOnVlan)

			intSetting.InterfaceName = fmt.Sprintf("port-channel %d", intSetting.PortChannelID)
			intSetting.Mode = constants.GetIntMode(obj.Edges.OnPoInterface.Edges.Mode)
			intSetting.VLANs = intVlan.String()
			intSetting.NativeVLAN = nVlan
			portChannelConfig = append(portChannelConfig, getInterfaceConfig(intSetting)...)
		} else {

			intVlan := types.GetHaveVlans(obj.Edges.HaveVlans)
			nVlan := types.GetNativeVlan(obj.Edges.NativeOnVlan)
			intSetting := &types.InterfaceSetting{
				InterfaceName: obj.InterfaceName,
				Mode:          mode,
				VLANs:         intVlan.String(),
				NativeVLAN:    nVlan,
				Shutdown:      obj.InterfaceShutdown,
				PoSetting:     false,
			}
			config = append(config, getInterfaceConfig(intSetting)...)
		}

	}
	config = append(config, portChannelConfig...)
	return
}
