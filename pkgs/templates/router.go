package templates

import (
	"fmt"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/mrzack99s/netcoco/pkgs/system"
)

func NewHTMLController(ft bool) {
	if ft {
		system.HtmlGroup = system.HttpRouter.Group(fmt.Sprintf("/%s/ui", system.Version))
		system.HttpRouter.LoadHTMLGlob(fmt.Sprintf("%sinitial.html", GetHTMLPath()))
		system.HttpRouter.Static("/js", fmt.Sprintf("%sjs", GetRootPath()))
		system.HttpRouter.Static("/css", fmt.Sprintf("%scss", GetRootPath()))
		system.HttpRouter.Static("/images", fmt.Sprintf("%simages", GetRootPath()))
		system.HtmlGroup.GET("initial", GetInitialPage)
		system.HttpRouter.GET("", func(c *gin.Context) {
			c.Request.URL.Path = fmt.Sprintf("/%s/ui/initial", system.Version)
			system.HttpRouter.HandleContext(c)
		})
	} else {
		// HTML
		system.HttpRouter.Static("/js", fmt.Sprintf("%sdist/js", GetRootPath()))
		system.HttpRouter.Static("/css", fmt.Sprintf("%sdist/css", GetRootPath()))
		system.HttpRouter.Static("/img", fmt.Sprintf("%sdist/img", GetRootPath()))
		system.HttpRouter.Static("/fonts", fmt.Sprintf("%sdist/fonts", GetRootPath()))
		system.HttpRouter.Use(static.Serve(fmt.Sprintf("/%s/ui", system.Version),
			static.LocalFile(fmt.Sprintf("%sdist", GetRootPath()), true)))
		system.HttpRouter.GET("", func(c *gin.Context) {
			c.Request.URL.Path = fmt.Sprintf("/%s/ui", system.Version)
			system.HttpRouter.HandleContext(c)
		})

		path := []string{"login", "device-setting", "dashboard", "devices", "device-type", "topology"}

		for _, str := range path {
			system.HttpRouter.Use(static.Serve(fmt.Sprintf("/%s/ui/%s", system.Version, str),
				static.LocalFile(fmt.Sprintf("%sdist", GetRootPath()), true)))
		}

	}
}
