package templates

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrzack99s/netcoco/pkgs/services"
)

func GetInitialPage(c *gin.Context) {
	c.HTML(http.StatusOK, "initial.html", gin.H{
		"title":           "Create a new administrator",
		"APIKeyTemporary": services.APIKeyTemporaryVar,
	})

}
