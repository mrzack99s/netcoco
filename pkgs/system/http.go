package system

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	ApplicationListener *http.Server
	HttpRouter          *gin.Engine
	SecureAPIGroup      *gin.RouterGroup
	UnsecureAPIGroup    *gin.RouterGroup
	HtmlGroup           *gin.RouterGroup
)
