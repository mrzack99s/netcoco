package services

import (
	"context"

	"github.com/mrzack99s/netcoco/ent"
)

func GetInterfaceLayer(client *ent.Client) (response []*ent.NetInterfaceLayer, err error) {
	response, err = client.NetInterfaceLayer.Query().All(context.Background())
	return
}
