package services

import (
	"context"

	"github.com/mrzack99s/netcoco/ent"
)

func GetDevicePlatform(client *ent.Client) (response []*ent.DevicePlatform, err error) {
	response, err = client.DevicePlatform.Query().All(context.Background())
	return
}
