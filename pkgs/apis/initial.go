package apis

import (
	"github.com/mrzack99s/netcoco/ent"
	"github.com/mrzack99s/netcoco/pkgs/security"
	"github.com/mrzack99s/netcoco/pkgs/services"
	"github.com/mrzack99s/netcoco/pkgs/system"
	"github.com/mrzack99s/netcoco/pkgs/templates"
)

func InitialSystem(client *ent.Client) {
	services.APIKeyTemporaryVar.Username = security.GeneratePasswordString(16)
	services.APIKeyTemporaryVar.APISecret = security.GeneratePasswordString(64)
	NewAdministratorController(system.SecureAPIGroup, client)
	templates.NewHTMLController(true)

}
