package apis

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mrzack99s/netcoco/ent"
	"github.com/mrzack99s/netcoco/pkgs/services"
)

type InterfaceController struct {
	client *ent.Client
	router gin.IRouter
}

func (ctl *InterfaceController) Create(c *gin.Context) {

	obj := ent.NetInterface{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "binding failed",
		})
		return
	}

	u, err := services.CreateInterface(ctl.client, obj)
	if err != nil {
		c.JSON(400, gin.H{

			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)

}

func (ctl *InterfaceController) CreateRange(c *gin.Context) {

	obj := []ent.NetInterface{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "binding failed",
		})
		return
	}

	u, err := services.CreateRangeInterface(ctl.client, obj)
	if err != nil {
		c.JSON(400, gin.H{

			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)

}

func (ctl *InterfaceController) Edit(c *gin.Context) {

	obj := ent.NetInterface{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "binding failed",
		})
		return
	}

	usr, err := services.EditInterfaceDetail(ctl.client, obj)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return

	} else {
		c.JSON(200, usr)

	}

}

func (ctl *InterfaceController) Get(c *gin.Context) {

	intId, _ := strconv.Atoi(c.Param("interfaceId"))

	usr, err := services.GetInterface(ctl.client, intId)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return

	} else {
		c.JSON(200, usr)

	}

}

func (ctl *InterfaceController) CleanInterface(c *gin.Context) {

	deviceId, _ := strconv.Atoi(c.Param("id"))

	err := services.CleanInterface(ctl.client, deviceId)
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

func (ctl *InterfaceController) Delete(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	err := services.DeleteInterface(ctl.client, id)
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

func NewInterfaceController(router gin.IRouter, client *ent.Client) *InterfaceController {
	pc := &InterfaceController{
		client: client,
		router: router,
	}

	pc.register()

	return pc

}

func (ctl *InterfaceController) register() {
	router := ctl.router.Group("/net-interface")
	router.POST("create", ctl.Create)
	router.POST("create-range", ctl.CreateRange)
	router.POST("edit", ctl.Edit)
	router.DELETE("delete/:id", ctl.Delete)
	router.DELETE("clean-interface/:id", ctl.CleanInterface)
	router.GET("get/:intId", ctl.Get)
}
