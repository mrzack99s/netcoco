package ios

import (
	"context"
	"fmt"
	"strings"

	"github.com/mrzack99s/netcoco/ent"
	"github.com/mrzack99s/netcoco/pkgs/netautomation/types"
)

func getIPRoutingConfig(o *types.IPStaticRoutingSetting) (config []string) {
	switch o.BrdInterface {
	case true:
		var rawCmd string = `ip route %s %s %s`
		replaceRawCmd := fmt.Sprintf(rawCmd, o.NetworkAddress, o.SubnetMask, o.InterfaceName)
		config = strings.Split(replaceRawCmd, "\n")

	case false:
		var rawCmd string = `ip route %s %s %s %s`
		replaceRawCmd := fmt.Sprintf(rawCmd, o.NetworkAddress, o.SubnetMask, o.InterfaceName, o.NextHop)
		config = strings.Split(replaceRawCmd, "\n")
	}

	return
}

func GetCommitIPStaticPoutingConfig(device *ent.Device) (config []string) {

	if i := device.QueryIPStaticRouting().CountX(context.Background()); i > 0 {
		device.Edges.IPStaticRouting = device.QueryIPStaticRouting().AllX(context.Background())
		for _, o := range device.Edges.IPStaticRouting {
			o.Edges.OnInterface = o.QueryOnInterface().OnlyX(context.Background())
			ipSetting := &types.IPStaticRoutingSetting{
				NetworkAddress: o.NetworkAddress,
				SubnetMask:     o.SubnetMask,
				NextHop:        o.NextHop,
				BrdInterface:   o.BrdInterface,
				InterfaceName:  o.Edges.OnInterface.InterfaceName,
			}
			config = append(config, getIPRoutingConfig(ipSetting)...)
		}
	}
	return
}
