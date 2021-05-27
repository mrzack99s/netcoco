package functions

import (
	"github.com/mrzack99s/netcoco/ent"
)

func FindDeleteVlanInVlanUsed(arr []*ent.Vlan, element *ent.DeletedVlanLog) bool {
	for _, o := range arr {
		if o.VlanID == element.VlanID {
			return true
		}
	}

	return false
}
