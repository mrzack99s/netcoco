package sg350

import (
	"context"
	"fmt"
	"strings"

	"github.com/mrzack99s/netcoco/ent"
	"github.com/mrzack99s/netcoco/pkgs/netautomation/types"
)

func getIPStaticRoutingConfig(o *types.IPStaticRoutingSetting) (config []string) {

	var rawCmd string = `ip route %s %s %s`
	replaceRawCmd := fmt.Sprintf(rawCmd, o.NetworkAddress, o.SubnetMask, o.InterfaceName)
	config = strings.Split(replaceRawCmd, "\n")

	return
}

func GetCommitIPStaticPoutingConfig(device *ent.Device) (config []string) {

	if i := device.QueryIPStaticRouting().CountX(context.Background()); i > 0 {
		device.Edges.IPStaticRouting = device.QueryIPStaticRouting().AllX(context.Background())
		for _, o := range device.Edges.IPStaticRouting {
			ipSetting := &types.IPStaticRoutingSetting{
				NetworkAddress: o.NetworkAddress,
				SubnetMask:     o.SubnetMask,
				NextHop:        o.NextHop,
			}
			config = append(config, getIPStaticRoutingConfig(ipSetting)...)
		}
	}
	return
}
