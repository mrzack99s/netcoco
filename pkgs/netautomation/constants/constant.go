package constants

import "github.com/mrzack99s/netcoco/ent"

const (
	MODE_ACCESS          = "access"
	MODE_TRUNKING        = "trunk"
	MODE_ETHERCHANNEL    = "etherchannel"
	MODE_NONE            = "none"
	DEVICE_TYPE_L3SWITCH = "l3switch"
	DEVICE_TYPE_L2SWITCH = "l2switch"
	DEVICE_TYPE_ROUTER   = "router"
)

func GetIntMode(intMode *ent.NetInterfaceMode) (mode string) {
	mode = MODE_NONE
	if intMode.InterfaceMode == "Access" {
		mode = MODE_ACCESS
	} else if intMode.InterfaceMode == "Trunking" {
		mode = MODE_TRUNKING
	} else if intMode.InterfaceMode == "EtherChannel" {
		mode = MODE_ETHERCHANNEL
	}
	return
}

func GetDeviceType(dType *ent.DeviceType) (mode string) {
	mode = MODE_NONE
	if dType.DeviceTypeName == "l3switch" {
		mode = DEVICE_TYPE_L3SWITCH
	} else if dType.DeviceTypeName == "l2switch" {
		mode = DEVICE_TYPE_L2SWITCH
	} else if dType.DeviceTypeName == "router" {
		mode = DEVICE_TYPE_ROUTER
	}
	return
}
