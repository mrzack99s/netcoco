package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/mrzack99s/netcoco/ent"
	"github.com/mrzack99s/netcoco/ent/deviceplatform"
	"github.com/mrzack99s/netcoco/ent/devicetype"
	"github.com/mrzack99s/netcoco/ent/netinterfacelayer"
	"github.com/mrzack99s/netcoco/ent/netinterfacemode"
	"github.com/mrzack99s/netcoco/pkgs/system"
)

func Open() (*ent.Client, error) {
	db, err := sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s",
			system.SystemConfigVar.NetCoCo.DB.SQL.Username,
			system.SystemConfigVar.NetCoCo.DB.SQL.Password,
			system.SystemConfigVar.NetCoCo.DB.SQL.Hostname,
			system.SystemConfigVar.NetCoCo.DB.SQL.DBName,
		),
	)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(25)
	db.SetMaxOpenConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)
	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB("mysql", db)
	return ent.NewClient(ent.Driver(drv)), nil
}

func DBInitial(client *ent.Client) {
	err := client.Schema.Create(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Device type initial
	if i := client.DeviceType.Query().CountX(context.Background()); i == 0 {
		temp := []string{"router", "l3switch", "l2switch", "firewall", "storage", "server"}
		for _, item := range temp {
			client.DeviceType.Create().SetDeviceTypeName(item).SaveX(context.Background())
		}
	} else {
		temp := []string{"router", "l3switch", "l2switch", "firewall", "storage", "server"}
		for _, item := range temp {
			if i := client.DeviceType.Query().Where(devicetype.DeviceTypeName(item)).CountX(context.Background()); i == 0 {
				client.DeviceType.Create().SetDeviceTypeName(item).SaveX(context.Background())
			}
		}
	}

	// Interface mode initial
	if i := client.NetInterfaceMode.Query().CountX(context.Background()); i == 0 {
		temp := []string{"Access", "Trunking", "EtherChannel", "None"}
		for _, item := range temp {
			client.NetInterfaceMode.Create().SetInterfaceMode(item).SaveX(context.Background())
		}
	} else {
		temp := []string{"Access", "Trunking", "EtherChannel", "None"}
		for _, item := range temp {
			if i := client.NetInterfaceMode.Query().Where(netinterfacemode.InterfaceModeEQ(item)).CountX(context.Background()); i == 0 {
				client.NetInterfaceMode.Create().SetInterfaceMode(item).SaveX(context.Background())
			}
		}
	}

	// Device platform initial
	if i := client.DevicePlatform.Query().CountX(context.Background()); i == 0 {
		temp := []string{"ios", "sg300", "sg350"}
		for _, item := range temp {
			client.DevicePlatform.Create().SetDevicePlatformName(item).SaveX(context.Background())
		}
	} else {
		temp := []string{"ios", "sg300", "sg350"}
		for _, item := range temp {
			if i := client.DevicePlatform.Query().Where(deviceplatform.DevicePlatformName(item)).CountX(context.Background()); i == 0 {
				client.DevicePlatform.Create().SetDevicePlatformName(item).SaveX(context.Background())
			}
		}

	}

	// Vlan1 initial
	if i := client.Vlan.Query().CountX(context.Background()); i == 0 {
		client.Vlan.Create().SetVlanID(1).SaveX(context.Background())
	}

	// Interface layer initial
	if i := client.NetInterfaceLayer.Query().CountX(context.Background()); i == 0 {
		for j := 2; j <= 3; j++ {
			client.NetInterfaceLayer.Create().SetInterfaceLayer(j).SaveX(context.Background())
		}
	} else {
		for j := 2; j <= 3; j++ {
			if i := client.NetInterfaceLayer.Query().Where(netinterfacelayer.InterfaceLayerEQ(j)).CountX(context.Background()); i == 0 {
				client.NetInterfaceLayer.Create().SetInterfaceLayer(j).SaveX(context.Background())
			}
		}

	}
}
