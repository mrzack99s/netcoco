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
		config = ios.GetCommitVlanConfig(device)
		config = append(config, ios.GetCommitInterfaceConfig(device)...)
	case "sg300":
		config = sg300.GetCommitVlanConfig(device)
		config = append(config, sg300.GetCommitInterfaceConfig(device)...)
	case "sg350":
		config = sg350.GetCommitVlanConfig(device)
		config = append(config, sg350.GetCommitInterfaceConfig(device)...)
	}
	return
}
