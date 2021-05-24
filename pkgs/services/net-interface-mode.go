package services

import (
	"context"

	"github.com/mrzack99s/netcoco/ent"
)

func GetInterfaceMode(client *ent.Client) (response []*ent.NetInterfaceMode, err error) {
	response, err = client.NetInterfaceMode.Query().All(context.Background())
	return
}
