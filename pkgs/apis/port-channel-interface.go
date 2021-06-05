package apis

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mrzack99s/netcoco/ent"
	"github.com/mrzack99s/netcoco/pkgs/services"
)

type PortChannelInterfaceController struct {
	client *ent.Client
	router gin.IRouter
}

func (ctl *PortChannelInterfaceController) Create(c *gin.Context) {

	obj := ent.PortChannelInterface{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "binding failed",
		})
		return
	}

	u, err := services.CreatePortChannelInterface(ctl.client, obj)
	if err != nil {
		c.JSON(400, gin.H{

			"error": err,
		})
		return
	}

	c.JSON(200, u)

}

func (ctl *PortChannelInterfaceController) Edit(c *gin.Context) {

	obj := ent.PortChannelInterface{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "binding failed",
		})
		return
	}

	usr, err := services.EditPortChannelInterface(ctl.client, obj)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return

	} else {
		c.JSON(200, usr)

	}

}

func (ctl *PortChannelInterfaceController) Get(c *gin.Context) {

	intId, _ := strconv.Atoi(c.Param("PortChannelInterfaceId"))

	usr, err := services.GetPortChannelInterface(ctl.client, intId)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return

	} else {
		c.JSON(200, usr)

	}

}

func (ctl *PortChannelInterfaceController) Delete(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	err := services.DeletePortChannelInterface(ctl.client, id)
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

func NewPortChannelInterfaceController(router gin.IRouter, client *ent.Client) *PortChannelInterfaceController {
	pc := &PortChannelInterfaceController{
		client: client,
		router: router,
	}

	pc.register()

	return pc

}

func (ctl *PortChannelInterfaceController) register() {
	router := ctl.router.Group("/po-interface")
	router.POST("create", ctl.Create)
	router.POST("edit", ctl.Edit)
	router.DELETE("delete/:id", ctl.Delete)
	router.GET("get/:intId", ctl.Get)
}
