package constants

import "github.com/mrzack99s/netcoco/ent"

const (
	MODE_ACCESS       = "access"
	MODE_TRUNKING     = "trunk"
	MODE_ETHERCHANNEL = "etherchannel"
	MODE_NONE         = "none"
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
