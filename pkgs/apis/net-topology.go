package apis

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mrzack99s/netcoco/ent"
	"github.com/mrzack99s/netcoco/pkgs/services"
)

type TopologyController struct {
	client *ent.Client
	router gin.IRouter
}

func (ctl *TopologyController) Create(c *gin.Context) {

	obj := ent.NetTopology{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "binding failed",
		})
		return
	}

	u, err := services.CreateTopology(ctl.client, obj)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(200, u)

}

func (ctl *TopologyController) Edit(c *gin.Context) {

	obj := ent.NetTopology{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "binding failed",
		})
		return
	}

	usr, err := services.EditTopologyDetail(ctl.client, obj)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return

	} else {
		c.JSON(200, usr)

	}

}

func (ctl *TopologyController) GetAll(c *gin.Context) {

	usr, err := services.GetAllTopology(ctl.client)
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

func (ctl *TopologyController) Get(c *gin.Context) {

	topoId, _ := strconv.Atoi(c.Param("topoId"))

	usr, err := services.GetTopology(ctl.client, topoId)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return

	} else {
		c.JSON(200, usr)

	}

}

func (ctl *TopologyController) Delete(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	err := services.DeleteTopology(ctl.client, id)
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

func NewTopologyController(router gin.IRouter, client *ent.Client) *TopologyController {
	pc := &TopologyController{
		client: client,
		router: router,
	}

	pc.register()

	return pc

}

func (ctl *TopologyController) register() {
	router := ctl.router.Group("/topology")
	router.POST("create", ctl.Create)
	router.POST("edit", ctl.Edit)
	router.DELETE("delete/:id", ctl.Delete)
	router.GET("get", ctl.GetAll)
	router.GET("get/:topoId", ctl.Get)
}
