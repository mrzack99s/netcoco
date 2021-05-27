package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/mrzack99s/netcoco/ent"
	"github.com/mrzack99s/netcoco/ent/administrator"
	"github.com/mrzack99s/netcoco/pkgs/security"
)

type ChangeAdminPassword struct {
	Username    string `json:"username"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

func GetAllAdministrator(client *ent.Client) (response []*ent.Administrator, err error) {
	response, err = client.Administrator.Query().All(context.Background())
	return
}

func CreateAdministrator(client *ent.Client, obj ent.Administrator) (response *ent.Administrator, err error) {
	response, err = client.Administrator.Create().
		SetUsername(obj.Username).
		SetPassword(security.Encrypt(obj.Password)).
		Save(context.Background())

	return
}

func ChangePasswordAdministrator(client *ent.Client, obj ChangeAdminPassword) (response *ent.Administrator, e error) {
	usr, err := client.Administrator.
		Query().Where(administrator.UsernameEQ(obj.Username)).
		Only(context.Background())
	if err != nil {
		response = nil
		e = errors.New(fmt.Sprintf("Not found username %s", obj.Username))
		return

	} else {
		if obj.OldPassword != security.Decrypt([]byte(usr.Password)) {
			response = nil
			e = errors.New("Password not match!")
			return

		} else {
			newPasswdEncrypt := security.Encrypt(obj.NewPassword)
			_, e = client.Administrator.UpdateOneID(usr.ID).
				SetUsername(usr.Username).
				SetPassword(newPasswdEncrypt).
				Save(context.Background())

			if e != nil {
				response = nil
				return
			}
			usr.Password = newPasswdEncrypt
			response = usr
			return

		}
	}
}

func CheckNilAdministrator(client *ent.Client) (status bool) {

	count, _ := client.Administrator.Query().Count(context.Background())
	status = true
	if count > 0 {
		status = false
	}
	return
}
