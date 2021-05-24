package templates

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrzack99s/netcoco/pkgs/services"
	"github.com/mrzack99s/netcoco/pkgs/system"
)

func GetLoginPage(c *gin.Context) {
	login_session, err := c.Cookie("logged")
	session := &services.LoggedSession{}
	authorized := false
	if err == nil {
		sessionCookie := services.LoggedSession{}
		json.Unmarshal([]byte(login_session), &sessionCookie)

		session, err = services.FindLoggedSession(sessionCookie.Username)
		if err == nil {
			services.CheckTimeout(session)
			authorized = session.Authorized
		}

	}
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title":         "Login to the system",
		"LoggedSession": session,
		"Authorized":    authorized,
		"SystemVersion": system.SystemConfigVar.MRZ_SW_AUTO.Version,
	})
}
