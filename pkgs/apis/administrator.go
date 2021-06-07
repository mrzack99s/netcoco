package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/mrzack99s/netcoco/ent"
	"github.com/mrzack99s/netcoco/pkgs/services"
)

type AdministratorController struct {
	client *ent.Client
	router gin.IRouter
}

func (ctl *AdministratorController) Create(c *gin.Context) {

	obj := ent.Administrator{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "binding failed",
		})
		return
	}

	u, err := services.CreateAdministrator(ctl.client, obj)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return

	} else {
		c.JSON(200, u)

	}

}

func (ctl *AdministratorController) Change(c *gin.Context) {

	obj := services.ChangeAdminPassword{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "binding failed",
		})
		return
	}

	usr, err := services.ChangePasswordAdministrator(ctl.client, obj)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return

	} else {
		c.JSON(200, usr)

	}

}

func NewAdministratorController(router gin.IRouter, client *ent.Client) *AdministratorController {
	pc := &AdministratorController{
		client: client,
		router: router,
	}

	pc.register()

	return pc

}

func (ctl *AdministratorController) register() {
	router := ctl.router.Group("/administrator")
	router.POST("create", ctl.Create)
	router.POST("change", ctl.Change)
}
