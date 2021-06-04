package ios

import (
	"context"
	"log"

	"github.com/mrzack99s/netcoco/pkgs/netautomation/types"
)

func GetCommands(task *types.Task) []string {
	return task.Commands
}
func SendConfig(task *types.Task) (err error) {

	device, err := task.Host.GetDriver()
	if err != nil {
		return
	}
	ctx, cancelOpen := context.WithCancel(context.Background())
	defer cancelOpen()
	//Connect
	err = device.Dial(ctx)
	if err != nil {
		return
	}
	defer device.Close(context.Background())

	//Enable
	err = device.Enable(context.Background())
	if err != nil {
		return
	}

	ctx, cancelRun := context.WithCancel(context.Background())
	defer cancelRun()

	_, err = device.Configure(ctx, task.Commands)
	if err != nil {
		return
	}

	_, err = device.Run(ctx, "write")
	if err != nil {
		return err
	}
	err = device.Close(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	err = nil
	return
}
