package netautomation

import (
	"github.com/mrzack99s/netcoco/ent"
	"github.com/mrzack99s/netcoco/pkgs/netautomation/ios"
	"github.com/mrzack99s/netcoco/pkgs/netautomation/sg300"
	"github.com/mrzack99s/netcoco/pkgs/netautomation/sg350"
)

func GetCommitConfig(device *ent.Device) (config []string) {
	switch device.Edges.InPlatform.DevicePlatformName {
	case "ios":
		if device.Edges.InType.DeviceTypeName != "router" {
			config = ios.GetCommitVlanConfig(device)
		}
		config = append(config, ios.GetCommitInterfaceConfig(device)...)
		config = append(config, ios.GetCommitIPStaticPoutingConfig(device)...)
	case "sg300":
		config = sg300.GetCommitVlanConfig(device)
		config = append(config, sg300.GetCommitInterfaceConfig(device)...)
	case "sg350":
		config = sg350.GetCommitVlanConfig(device)
		config = append(config, sg350.GetCommitInterfaceConfig(device)...)
		config = append(config, sg350.GetCommitIPStaticPoutingConfig(device)...)
	}

	switch device.Edges.InType.DeviceTypeName {
	case "l3switch":
		config = append(config, "ip routing")
	}
	return
}
