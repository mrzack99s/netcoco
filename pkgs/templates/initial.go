package templates

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrzack99s/netcoco/pkgs/services"
	"github.com/mrzack99s/netcoco/pkgs/system"
)

func GetInitialPage(c *gin.Context) {
	c.HTML(http.StatusOK, "initial.html", gin.H{
		"title":           "Create a new administrator",
		"SystemVersion":   system.SystemConfigVar.MRZ_SW_AUTO.Version,
		"APIKeyTemporary": services.APIKeyTemporaryVar,
	})

}
