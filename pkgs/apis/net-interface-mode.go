package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/mrzack99s/netcoco/ent"
	"github.com/mrzack99s/netcoco/pkgs/services"
)

type InterfaceModeController struct {
	client *ent.Client
	router gin.IRouter
}

func (ctl *InterfaceModeController) GetAll(c *gin.Context) {

	usr, err := services.GetInterfaceMode(ctl.client)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return

	} else {
		c.JSON(200, usr)

	}

}

func NewInterfaceModeController(router gin.IRouter, client *ent.Client) *InterfaceModeController {
	pc := &InterfaceModeController{
		client: client,
		router: router,
	}

	pc.register()

	return pc

}

func (ctl *InterfaceModeController) register() {
	router := ctl.router.Group("/net-interface-mode")
	router.GET("get", ctl.GetAll)
}
