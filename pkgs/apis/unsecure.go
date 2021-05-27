package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/mrzack99s/netcoco/ent"
	"github.com/mrzack99s/netcoco/pkgs/services"
)

type UnsecureController struct {
	client *ent.Client
	router gin.IRouter
}

func (ctl *UnsecureController) Login(c *gin.Context) {

	obj := ent.Administrator{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "binding failed",
		})
		return
	}

	usr, err := services.GetAuthentication(ctl.client, obj)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return

	} else {
		c.JSON(200, usr)

	}
}

func (ctl *UnsecureController) CheckSession(c *gin.Context) {

	obj := ent.Administrator{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "binding failed",
		})
		return
	}

	usr, err := services.CheckSessionWithUsernameAndAPISecret(obj)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return

	} else {
		c.JSON(200, usr)

	}
}

func (ctl *UnsecureController) Logout(c *gin.Context) {

	obj := ent.Administrator{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "binding failed",
		})
		return
	}

	session, err := services.FindLoggedSession(obj.Username)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return

	} else {
		services.Logout(session)
		c.JSON(200, gin.H{
			"status": true,
		})
	}
}

func NewUnsecureController(router gin.IRouter, client *ent.Client) *UnsecureController {
	pc := &UnsecureController{
		client: client,
		router: router,
	}

	pc.register()

	return pc

}

func (ctl *UnsecureController) register() {
	ctl.router.POST("login", ctl.Login)
	ctl.router.POST("logout", ctl.Logout)
	ctl.router.POST("check", ctl.CheckSession)
}
