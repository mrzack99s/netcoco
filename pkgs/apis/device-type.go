package apis

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mrzack99s/netcoco/ent"
	"github.com/mrzack99s/netcoco/pkgs/services"
)

type VlanController struct {
	client *ent.Client
	router gin.IRouter
}

func (ctl *VlanController) Create(c *gin.Context) {

	obj := ent.Vlan{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "binding failed",
		})
		return
	}

	u, err := services.CreateVlan(ctl.client, obj)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)

}

func (ctl *VlanController) Edit(c *gin.Context) {

	obj := ent.Vlan{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "binding failed",
		})
		return
	}

	usr, err := services.EditVlanDetail(ctl.client, obj)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return

	} else {
		c.JSON(200, usr)

	}

}

func (ctl *VlanController) GetAll(c *gin.Context) {

	usr, err := services.GetAllVlan(ctl.client)
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

func (ctl *VlanController) Get(c *gin.Context) {

	vlanId, _ := strconv.Atoi(c.Param("vlanId"))

	usr, err := services.GetVlan(ctl.client, vlanId)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return

	} else {
		c.JSON(200, usr)

	}

}

func (ctl *VlanController) Delete(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	err := services.DeleteVlan(ctl.client, id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return

	} else {
		c.JSON(200, gin.H{
			"status": true,
		})

	}

}

func NewVlanController(router gin.IRouter, client *ent.Client) *VlanController {
	pc := &VlanController{
		client: client,
		router: router,
	}

	pc.register()

	return pc

}

func (ctl *VlanController) register() {
	router := ctl.router.Group("/vlan")
	router.POST("create", ctl.Create)
	router.POST("edit", ctl.Edit)
	router.DELETE("delete/:id", ctl.Delete)
	router.GET("get", ctl.GetAll)
	router.GET("get/:vlanId", ctl.Get)
}
