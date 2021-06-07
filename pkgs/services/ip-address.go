package services

import (
	"context"
	"fmt"
	"net"

	"github.com/mrzack99s/netcoco/ent"
	"github.com/mrzack99s/netcoco/ent/device"
	"github.com/mrzack99s/netcoco/ent/ipaddress"
)

func CheckNotDuplicateIPAddress(client *ent.Client, d *ent.Device, ip_address *ent.IPAddress) (*ent.IPAddress, error) {
	found, err := client.IPAddress.Query().
		Where(
			ipaddress.And(
				ipaddress.IPAddressEQ(ip_address.IPAddress),
				ipaddress.SubnetMaskEQ(ip_address.SubnetMask),
				ipaddress.HasOnDeviceWith(device.IDEQ(d.ID)),
			),
		).Only(context.Background())
	if err != nil {
		return nil, err
	}
	return found, nil

}

func CheckNetworkOverlap(client *ent.Client, d *ent.Device, ip_address *ent.IPAddress) bool {
	allInterface := client.IPAddress.Query().Where(
		ipaddress.HasOnDeviceWith(device.IDEQ(d.ID)),
	).AllX(context.Background())
	for _, ip := range allInterface {
		if ip.SubnetMask != ip_address.SubnetMask {
			if ipIntersect(ip, ip_address) {
				return true
			}
		}
	}

	return false

}

func ipIntersect(n1, n2 *ent.IPAddress) bool {
	n1Mask := net.IPMask(net.ParseIP(n1.SubnetMask).To4())
	n2Mask := net.IPMask(net.ParseIP(n2.SubnetMask).To4())
	n1PrefixSize, _ := n1Mask.Size()
	n2PrefixSize, _ := n2Mask.Size()
	_, n1IpCIDR, _ := net.ParseCIDR(fmt.Sprintf("%s/%d", n1.IPAddress, n1PrefixSize))
	_, n2IpCIDR, _ := net.ParseCIDR(fmt.Sprintf("%s/%d", n2.IPAddress, n2PrefixSize))
	return n2IpCIDR.Contains(n1IpCIDR.IP) || n1IpCIDR.Contains(n2IpCIDR.IP)
}
