package apis

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mrzack99s/netcoco/ent"
	"github.com/mrzack99s/netcoco/pkgs/services"
)

type DeviceTypeController struct {
	client *ent.Client
	router gin.IRouter
}

func (ctl *DeviceTypeController) GetAll(c *gin.Context) {

	usr, err := services.GetAllDeviceType(ctl.client)
	if err != nil {
		c.JSON(400, gin.H{
			"status": false,
			"error":  err,
		})
		return

	} else {
		c.JSON(200, usr)
	}

}

func (ctl *DeviceTypeController) Get(c *gin.Context) {

	deviceTypeId, _ := strconv.Atoi(c.Param("deviceTypeId"))

	usr, err := services.GetDeviceType(ctl.client, deviceTypeId)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return

	} else {
		c.JSON(200, usr)

	}

}

func NewDeviceTypeController(router gin.IRouter, client *ent.Client) *DeviceTypeController {
	pc := &DeviceTypeController{
		client: client,
		router: router,
	}

	pc.register()

	return pc

}

func (ctl *DeviceTypeController) register() {
	router := ctl.router.Group("/device-type")
	router.GET("get", ctl.GetAll)
	router.GET("get/:deviceTypeId", ctl.Get)
}
