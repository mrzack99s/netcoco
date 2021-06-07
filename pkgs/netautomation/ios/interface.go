package ios

import (
	"context"
	"fmt"
	"strings"

	"github.com/mrzack99s/netcoco/ent"
	"github.com/mrzack99s/netcoco/pkgs/netautomation/constants"
	"github.com/mrzack99s/netcoco/pkgs/netautomation/types"
)

func getInterfaceL3Config(o *types.InterfaceL3Setting) (config []string) {
	switch o.Shutdown {
	case true:
		switch o.DeviceType {
		case constants.DEVICE_TYPE_L3SWITCH:
			var rawCmd string = `interface %s
no switchport
ip address %s %s
shutdown
exit`
			replaceRawCmd := fmt.Sprintf(rawCmd, o.InterfaceName, o.IPAddress, o.SubnetMask)
			config = strings.Split(replaceRawCmd, "\n")
		case constants.DEVICE_TYPE_ROUTER:
			var rawCmd string = `interface %s
ip address %s %s
shutdown
exit`
			replaceRawCmd := fmt.Sprintf(rawCmd, o.InterfaceName, o.IPAddress, o.SubnetMask)
			config = strings.Split(replaceRawCmd, "\n")
		}

	case false:
		switch o.DeviceType {
		case constants.DEVICE_TYPE_L3SWITCH:
			var rawCmd string = `interface %s
no switchport
ip address %s %s
no shutdown
exit`
			replaceRawCmd := fmt.Sprintf(rawCmd, o.InterfaceName, o.IPAddress, o.SubnetMask)
			config = strings.Split(replaceRawCmd, "\n")
		case constants.DEVICE_TYPE_ROUTER:
			var rawCmd string = `interface %s
ip address %s %s
no shutdown
exit`
			replaceRawCmd := fmt.Sprintf(rawCmd, o.InterfaceName, o.IPAddress, o.SubnetMask)
			config = strings.Split(replaceRawCmd, "\n")
		}
	}
	return
}

func getInterfaceL2Config(o *types.InterfaceL2Setting) (config []string) {

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
	var haveEtherChannel bool = false
	for _, obj := range device.Edges.Interfaces {
		obj.Edges.OnLayer = obj.QueryOnLayer().OnlyX(context.Background())
		switch obj.Edges.OnLayer.InterfaceLayer {
		case 2:
			obj.Edges.Mode = obj.QueryMode().OnlyX(context.Background())
			mode := constants.GetIntMode(obj.Edges.Mode)

			if mode == constants.MODE_ETHERCHANNEL {
				obj.Edges.OnPoInterface = obj.QueryOnPoInterface().OnlyX(context.Background())
				obj.Edges.OnPoInterface.Edges.Mode = obj.Edges.OnPoInterface.QueryMode().OnlyX(context.Background())
				if i := obj.Edges.OnPoInterface.QueryHaveVlans().CountX(context.Background()); i > 0 {
					obj.Edges.OnPoInterface.Edges.HaveVlans = obj.Edges.OnPoInterface.QueryHaveVlans().AllX(context.Background())
				}
				if i := obj.Edges.OnPoInterface.QueryNativeOnVlan().CountX(context.Background()); i > 0 {
					obj.Edges.OnPoInterface.Edges.NativeOnVlan = obj.Edges.OnPoInterface.QueryNativeOnVlan().OnlyX(context.Background())
				}
			} else {
				obj.Edges.HaveVlans = obj.QueryHaveVlans().AllX(context.Background())
				if i := obj.QueryNativeOnVlan().CountX(context.Background()); i > 0 {
					obj.Edges.NativeOnVlan = obj.QueryNativeOnVlan().OnlyX(context.Background())
				}

			}

			if mode == constants.MODE_ETHERCHANNEL {
				intSetting := &types.InterfaceL2Setting{
					InterfaceName: obj.InterfaceName,
					PortChannelID: obj.Edges.OnPoInterface.PoInterfaceID,
					Mode:          mode,
					PoSetting:     true,
					Shutdown:      obj.InterfaceShutdown,
				}
				config = append(config, getInterfaceL2Config(intSetting)...)
				haveEtherChannel = true

			} else {

				intVlan := types.GetHaveVlans(obj.Edges.HaveVlans)
				nVlan := types.GetNativeVlan(obj.Edges.NativeOnVlan)
				intSetting := &types.InterfaceL2Setting{
					InterfaceName: obj.InterfaceName,
					Mode:          mode,
					VLANs:         intVlan.String(),
					NativeVLAN:    nVlan,
					Shutdown:      obj.InterfaceShutdown,
					PoSetting:     false,
				}
				config = append(config, getInterfaceL2Config(intSetting)...)
			}
		case 3:
			obj.Edges.OnIPAddress = obj.QueryOnIPAddress().OnlyX(context.Background())
			obj.Edges.OnDevice = obj.QueryOnDevice().OnlyX(context.Background())
			obj.Edges.OnDevice.Edges.InType = obj.QueryOnDevice().QueryInType().OnlyX(context.Background())
			intSetting := &types.InterfaceL3Setting{
				InterfaceName: obj.InterfaceName,
				Shutdown:      obj.InterfaceShutdown,
				IPAddress:     obj.Edges.OnIPAddress.IPAddress,
				SubnetMask:    obj.Edges.OnIPAddress.SubnetMask,
				DeviceType:    constants.GetDeviceType(obj.Edges.OnDevice.Edges.InType),
			}
			config = append(config, getInterfaceL3Config(intSetting)...)
		}

	}
	if haveEtherChannel {
		device.Edges.PoInterfaces = device.QueryPoInterfaces().WithOnLayer().AllX(context.Background())
		for _, po := range device.Edges.PoInterfaces {
			switch po.Edges.OnLayer.InterfaceLayer {
			case 2:
				po.Edges.Mode = po.QueryMode().OnlyX(context.Background())
				if i := po.QueryHaveVlans().CountX(context.Background()); i > 0 {
					po.Edges.HaveVlans = po.QueryHaveVlans().AllX(context.Background())
				}
				if i := po.QueryNativeOnVlan().CountX(context.Background()); i > 0 {
					po.Edges.NativeOnVlan = po.QueryNativeOnVlan().OnlyX(context.Background())
				}
				intVlan := types.GetHaveVlans(po.Edges.HaveVlans)
				nVlan := types.GetNativeVlan(po.Edges.NativeOnVlan)
				intSetting := &types.InterfaceL2Setting{
					InterfaceName: fmt.Sprintf("port-channel %d", po.ID),
					Mode:          po.Edges.Mode.InterfaceMode,
					VLANs:         intVlan.String(),
					NativeVLAN:    nVlan,
					Shutdown:      po.PoInterfaceShutdown,
					PoSetting:     true,
				}
				config = append(config, getInterfaceL2Config(intSetting)...)
			case 3:
				po.Edges.OnIPAddress = po.QueryOnIPAddress().OnlyX(context.Background())
				po.Edges.OnDevice = po.QueryOnDevice().OnlyX(context.Background())
				po.Edges.OnDevice.Edges.InType = po.QueryOnDevice().QueryInType().OnlyX(context.Background())
				intSetting := &types.InterfaceL3Setting{
					InterfaceName: fmt.Sprintf("port-channel %d", po.ID),
					Shutdown:      po.PoInterfaceShutdown,
					IPAddress:     po.Edges.OnIPAddress.IPAddress,
					SubnetMask:    po.Edges.OnIPAddress.SubnetMask,
					DeviceType:    constants.GetDeviceType(po.Edges.OnDevice.Edges.InType),
				}
				config = append(config, getInterfaceL3Config(intSetting)...)
			}
		}
	}
	return
}
