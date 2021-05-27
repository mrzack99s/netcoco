package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	entsql "entgo.io/ent/dialect/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	cors "github.com/itsjamie/gin-cors"
	"github.com/mrzack99s/netcoco/ent"
	"github.com/mrzack99s/netcoco/pkgs/apis"
	"github.com/mrzack99s/netcoco/pkgs/services"
	"github.com/mrzack99s/netcoco/pkgs/system"
)

var (
	product_mode = "development"
	version      = "v1beta"
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
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB("mysql", db)
	return ent.NewClient(ent.Driver(drv)), nil
}

func SystemInitial(client *ent.Client) {
	system.SystemVersion = version
	mode := gin.DebugMode
	if product_mode == "production" {
		mode = gin.ReleaseMode
	}

	gin.SetMode(mode)

	// API
	system.HttpRouter = gin.Default()
	system.HttpRouter.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	//Unsecure API
	system.UnsecureAPIGroup = system.HttpRouter.Group(fmt.Sprintf("/%s/unsecure/api", system.SystemVersion))
	//Secure API
	system.SecureAPIGroup = system.HttpRouter.Group(fmt.Sprintf("/%s/api", system.SystemVersion), services.APIAuthentication)
	apis.NewUnsecureController(system.UnsecureAPIGroup, client)

	system.ApplicationListener = &http.Server{
		Addr:    ":8080",
		Handler: system.HttpRouter,
	}
}

func main() {
	system.ParseSystemConfig("./config.yaml")

	client, err := Open()
	if err != nil {
		log.Fatal(err)
	}

	client.Schema.Create(context.Background())
	if i, _ := client.DeviceType.Query().Count(context.Background()); i == 0 {
		temp := []string{"router", "l3switch", "l2switch", "firewall", "storage", "server"}
		for _, item := range temp {
			client.DeviceType.Create().SetDeviceTypeName(item).Save(context.Background())
		}
	}

	if i, _ := client.NetInterfaceMode.Query().Count(context.Background()); i == 0 {
		temp := []string{"Access", "Trunking", "None"}
		for _, item := range temp {
			client.NetInterfaceMode.Create().SetInterfaceMode(item).Save(context.Background())
		}
	}

	if i, _ := client.DevicePlatform.Query().Count(context.Background()); i == 0 {
		temp := []string{"ios", "sg300", "sg350"}
		for _, item := range temp {
			client.DevicePlatform.Create().SetDevicePlatformName(item).Save(context.Background())
		}
	}
	if i, _ := client.Vlan.Query().Count(context.Background()); i == 0 {
		client.Vlan.Create().SetVlanID(1).Save(context.Background())

	}

	SystemInitial(client)

	// API Initial
	count := services.CheckNilAdministrator(client)
	if count {
		apis.InitialSystem(client)
		ctx, cancelCtx := context.WithCancel(context.Background())
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				default:
					system.ApplicationListener.ListenAndServe()
				}
			}
		}()

		for {
			if !services.CheckNilAdministrator(client) {
				cancelCtx()
				system.HttpRouter = nil
				system.ApplicationListener.Shutdown(context.Background())
				break
			}
		}

		SystemInitial(client)
		apis.DefaultSystem(client)
		system.ApplicationListener.ListenAndServe()

	} else {
		apis.DefaultSystem(client)
		system.ApplicationListener.ListenAndServe()
	}
}
