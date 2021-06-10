package apis

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mrzack99s/netcoco/ent"
	"github.com/mrzack99s/netcoco/pkgs/services"
)

type IPRoutingController struct {
	client *ent.Client
	router gin.IRouter
}

func (ctl *IPRoutingController) Create(c *gin.Context) {

	obj := ent.IPStaticRoutingTable{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "binding failed",
		})
		return
	}

	u, err := services.CreateIPRouting(ctl.client, obj)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)

}

func (ctl *IPRoutingController) GetConnected(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("did"))

	u, err := services.GetCanBeStaticRouteOnDevice(ctl.client, id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)

}

func (ctl *IPRoutingController) Delete(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	err := services.DeleteIPRouting(ctl.client, id)
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

func NewIPRoutingController(router gin.IRouter, client *ent.Client) *IPRoutingController {
	pc := &IPRoutingController{
		client: client,
		router: router,
	}

	pc.register()

	return pc

}

func (ctl *IPRoutingController) register() {
	router := ctl.router.Group("/ip-static-routing-table")
	router.POST("create", ctl.Create)
	router.GET("get/connected/:did", ctl.GetConnected)
	router.DELETE("delete/:id", ctl.Delete)
}
