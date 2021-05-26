package apis

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mrzack99s/netcoco/ent"
	"github.com/mrzack99s/netcoco/pkgs/services"
)

type DeviceController struct {
	client *ent.Client
	router gin.IRouter
}

func (ctl *DeviceController) Create(c *gin.Context) {

	obj := ent.Device{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "binding failed",
		})
		return
	}

	u, err := services.CreateDevice(ctl.client, obj)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(200, u)

}

func (ctl *DeviceController) AddDeviceVlan(c *gin.Context) {

	obj := ent.Device{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "binding failed",
		})
		return
	}

	u, err := services.AddDeviceVlan(ctl.client, obj)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(200, u)

}

func (ctl *DeviceController) Edit(c *gin.Context) {

	obj := ent.Device{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "binding failed",
		})
		return
	}

	usr, err := services.EditDeviceDetail(ctl.client, obj)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return

	} else {
		c.JSON(200, usr)

	}

}

func (ctl *DeviceController) GetAll(c *gin.Context) {

	usr, err := services.GetAllDevice(ctl.client)
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

func (ctl *DeviceController) Get(c *gin.Context) {

	deviceId, _ := strconv.Atoi(c.Param("deviceId"))

	usr, err := services.GetDevice(ctl.client, deviceId)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return

	} else {
		c.JSON(200, usr)

	}

}

func (ctl *DeviceController) Delete(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	err := services.DeleteDevice(ctl.client, id)
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

func (ctl *DeviceController) DeleteDeviceVlan(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	vid, _ := strconv.Atoi(c.Param("vid"))

	err := services.DeleteDeviceVlan(ctl.client, id, vid)
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

func NewDeviceController(router gin.IRouter, client *ent.Client) *DeviceController {
	pc := &DeviceController{
		client: client,
		router: router,
	}

	pc.register()

	return pc

}

func (ctl *DeviceController) register() {
	router := ctl.router.Group("/device")
	router.POST("create", ctl.Create)
	router.POST("add-vlan", ctl.AddDeviceVlan)
	router.POST("edit", ctl.Edit)
	router.DELETE("delete/:id", ctl.Delete)
	router.DELETE("delete-vlan/:id/:vid", ctl.DeleteDeviceVlan)
	router.GET("get", ctl.GetAll)
	router.GET("get/:deviceId", ctl.Get)
}
