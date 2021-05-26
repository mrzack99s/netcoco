package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/mrzack99s/netcoco/ent"
	"github.com/mrzack99s/netcoco/pkgs/services"
)

type DevicePlatformController struct {
	client *ent.Client
	router gin.IRouter
}

func (ctl *DevicePlatformController) GetAll(c *gin.Context) {

	usr, err := services.GetDevicePlatform(ctl.client)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return

	} else {
		c.JSON(200, usr)

	}

}

func NewDevicePlatformController(router gin.IRouter, client *ent.Client) *DevicePlatformController {
	pc := &DevicePlatformController{
		client: client,
		router: router,
	}

	pc.register()

	return pc

}

func (ctl *DevicePlatformController) register() {
	router := ctl.router.Group("/device-platform")
	router.GET("get", ctl.GetAll)
}
