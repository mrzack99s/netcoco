package apis

import (
	"github.com/mrzack99s/netcoco/ent"
	"github.com/mrzack99s/netcoco/pkgs/system"
	"github.com/mrzack99s/netcoco/pkgs/templates"
)

func DefaultSystem(client *ent.Client) {
	NewAdministratorController(system.SecureAPIGroup, client)
	NewDeviceController(system.SecureAPIGroup, client)
	NewInterfaceController(system.SecureAPIGroup, client)
	NewTopologyController(system.SecureAPIGroup, client)
	NewTopologyDeviceMapController(system.SecureAPIGroup, client)
	NewDeviceTypeController(system.SecureAPIGroup, client)
	NewInterfaceModeController(system.SecureAPIGroup, client)
	NewNetAutomationController(system.SecureAPIGroup, client)
	NewDevicePlatformController(system.SecureAPIGroup, client)
	NewVlanController(system.SecureAPIGroup, client)
	templates.NewHTMLController(false)
}
