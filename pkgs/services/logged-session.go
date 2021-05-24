package services

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mrzack99s/netcoco/ent"
	"github.com/mrzack99s/netcoco/ent/administrator"
	"github.com/mrzack99s/netcoco/pkgs/security"
)

type LoggedSession struct {
	Username   string    `json:"username"`
	APISecret  string    `json:"api_secret"`
	Authorized bool      `json:"authorized"`
	AtDatetime time.Time `json:"at_datetime"`
	Reason     string    `json:"reason"`
}

type APIKeyTemporary struct {
	Username  string `json:"username"`
	APISecret string `json:"api_secret"`
}

var LoggedSessionVar = make(map[string]*LoggedSession)
var APIKeyTemporaryVar = &APIKeyTemporary{}

func FindLoggedSession(username string) (session *LoggedSession, err error) {
	if v, found := LoggedSessionVar[username]; found {
		session = v
		return
	}
	err = errors.New("Not found session")
	return
}

func CheckTimeout(session *LoggedSession) {
	nowTime := time.Now()
	diffTime := nowTime.Sub(session.AtDatetime)
	if int(diffTime.Hours()) > 0 && int(diffTime.Minutes()) > 45 {
		session.Authorized = false
		session.Reason = "time-out"
	}

}

func Logout(session *LoggedSession) {
	delete(LoggedSessionVar, session.Username)
}

func CheckSessionWithUsernameAndAPISecret(obj ent.Administrator) (auth *LoggedSession, e error) {
	session, err := FindLoggedSession(obj.Username)
	if err != nil {
		e = errors.New("Not found this session ")
		return
	} else {
		if session.APISecret == obj.Password {
			auth = session
		} else {
			e = errors.New("API Secret not match!")
		}
	}
	return
}

func GetAuthentication(client *ent.Client, obj ent.Administrator) (auth *LoggedSession, e error) {
	response, err := client.Administrator.Query().Where(administrator.UsernameEQ(obj.Username)).Only(context.Background())
	if err != nil {
		e = err
		return
	} else {
		if security.Decrypt([]byte(response.Password)) == obj.Password {
			session, err := FindLoggedSession(response.Username)
			if err == nil && session.Authorized {
				auth = session
			} else {
				auth = &LoggedSession{
					Username:   response.Username,
					APISecret:  security.GeneratePasswordString(128),
					Authorized: true,
					AtDatetime: time.Now(),
				}
				LoggedSessionVar[response.Username] = auth
			}
		} else {
			e = errors.New("Password is not correct!")
			return
		}

	}
	return
}

func APIAuthentication(c *gin.Context) {
	username, apiSecret, hasAuth := c.Request.BasicAuth()
	authorized := false

	if hasAuth {
		if APIKeyTemporaryVar != nil && username == APIKeyTemporaryVar.Username && apiSecret == APIKeyTemporaryVar.APISecret {
			authorized = true
		} else {
			session, err := FindLoggedSession(username)
			CheckTimeout(session)
			if err == nil {
				if session.APISecret == apiSecret {
					authorized = true
				}
			}
		}
	}

	if authorized {
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

}
