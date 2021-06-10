package types

type IPStaticRoutingSetting struct {
	NetworkAddress string
	SubnetMask     string
	NextHop        string
	InterfaceName  string
	BrdInterface   bool
}
