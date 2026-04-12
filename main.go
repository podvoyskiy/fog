package main

import (
	"os"

	"github.com/podvoyskiy/fog/cmds"
	"github.com/podvoyskiy/fog/config"
	"github.com/podvoyskiy/fog/ui"
	u "github.com/podvoyskiy/fog/utils"
)

func main() {
	configDir := u.Must(os.UserConfigDir())
	config := u.Must(config.Load(configDir))

	args := os.Args[1:]

	switch len(args) {
	case 0:
		ui.Run(config)
		break
	default:
		if err := cmds.HandleCmd(config, args); err != nil {
			u.Yellow().Println(err)
			return
		}
		break
	}
}
