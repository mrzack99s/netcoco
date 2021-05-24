package apis

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mrzack99s/netcoco/ent"
	"github.com/mrzack99s/netcoco/pkgs/services"
)

type TopologyDeviceMapController struct {
	client *ent.Client
	router gin.IRouter
}

func (ctl *TopologyDeviceMapController) Create(c *gin.Context) {

	obj := services.DeviceMapStruct{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "binding failed",
		})
		return
	}

	u, err := services.CreateDeviceMap(ctl.client, obj)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(200, u)

}

func (ctl *TopologyDeviceMapController) CreateEdge(c *gin.Context) {

	obj := services.EdgeStruct{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "binding failed",
		})
		return
	}

	u, err := services.CreateEdge(ctl.client, obj)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(200, u)

}

func (ctl *TopologyDeviceMapController) Delete(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	err := services.DeleteDeviceMap(ctl.client, id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return

	} else {
		c.JSON(200, gin.H{
			"status": true,
		})

	}

}

func NewTopologyDeviceMapController(router gin.IRouter, client *ent.Client) *TopologyDeviceMapController {
	pc := &TopologyDeviceMapController{
		client: client,
		router: router,
	}

	pc.register()

	return pc

}

func (ctl *TopologyDeviceMapController) register() {
	router := ctl.router.Group("/topology-device-map")
	router.POST("create", ctl.Create)
	router.POST("create-edge", ctl.CreateEdge)
	router.DELETE("delete/:id", ctl.Delete)

}
