package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/mrzack99s/netcoco/ent"
	"github.com/mrzack99s/netcoco/pkgs/services"
)

type InterfaceLayerController struct {
	client *ent.Client
	router gin.IRouter
}

func (ctl *InterfaceLayerController) GetAll(c *gin.Context) {

	usr, err := services.GetInterfaceLayer(ctl.client)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return

	} else {
		c.JSON(200, usr)

	}

}

func NewInterfaceLayerController(router gin.IRouter, client *ent.Client) *InterfaceLayerController {
	pc := &InterfaceLayerController{
		client: client,
		router: router,
	}

	pc.register()

	return pc

}

func (ctl *InterfaceLayerController) register() {
	router := ctl.router.Group("/net-interface-layer")
	router.GET("get", ctl.GetAll)
}
