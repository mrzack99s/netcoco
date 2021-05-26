package apis

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mrzack99s/netcoco/ent"
	"github.com/mrzack99s/netcoco/pkgs/services"
)

type NetAutomationController struct {
	client *ent.Client
	router gin.IRouter
}

func (ctl *NetAutomationController) GetPingTopology(c *gin.Context) {

	topoId, _ := strconv.Atoi(c.Param("topoId"))

	usr, err := services.GetPingTopology(ctl.client, topoId)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return

	} else {
		c.JSON(200, usr)

	}

}

func (ctl *NetAutomationController) GetPingDevice(c *gin.Context) {

	deviceId, _ := strconv.Atoi(c.Param("deviceId"))

	usr, err := services.GetPingDevice(ctl.client, deviceId)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return

	} else {
		c.JSON(200, usr)

	}

}

func (ctl *NetAutomationController) Commit(c *gin.Context) {

	obj := ent.Device{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "binding failed",
		})
		return
	}

	usr, err := services.Commit(ctl.client, obj.ID)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return

	} else {
		c.JSON(200, usr)

	}

}

func NewNetAutomationController(router gin.IRouter, client *ent.Client) *NetAutomationController {
	pc := &NetAutomationController{
		client: client,
		router: router,
	}

	pc.register()

	return pc

}

func (ctl *NetAutomationController) register() {
	router := ctl.router.Group("/net-automation")
	router.POST("commit", ctl.Commit)
	router.GET("ping-topo/:topoId", ctl.GetPingTopology)
	router.GET("ping-device/:deviceId", ctl.GetPingDevice)

}
