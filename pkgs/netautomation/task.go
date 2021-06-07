package netautomation

import (
	"fmt"

	"github.com/mrzack99s/netcoco/pkgs/netautomation/types"
)

type Task types.Task

func (task *Task) GetCommands() []string {
	return task.Commands
}
func (task *Task) SendConfig() (err error) {

	fmt.Println(task.Commands)

	// switch task.Host.Platform {
	// case "ios":
	// 	err = ios.SendConfig((*types.Task)(task))
	// case "sg300":
	// 	err = sg300.SendConfig((*types.Task)(task))
	// case "sg350":
	// 	err = sg350.SendConfig((*types.Task)(task))
	// }
	return
}
