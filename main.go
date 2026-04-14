package main

import (
	"fmt"
	"os"

	"github.com/podvoyskiy/fog/cmds"
	"github.com/podvoyskiy/fog/config"
	"github.com/podvoyskiy/fog/ui"
	u "github.com/podvoyskiy/fog/utils"
)

var (
	version   = "dev"
	buildTime = "unknown"
)

func main() {
	args := os.Args[1:]

	if len(args) == 1 && (args[0] == "--version" || args[0] == "-v") {
		u.Cyan().Printf("fog %s (built at %s)\n", version, buildTime)
		return
	}

	configDir := u.Must(os.UserConfigDir())
	config := u.Must(config.Load(configDir))

	switch len(args) {
	case 0:
		if err := ui.Run(config); err != nil {
			//print error as echo command to prevent eval from executing it
			fmt.Printf("echo '%s'\n", u.Red().Sprint(err))
			os.Exit(1)
		}
	default:
		if err := cmds.HandleCmd(config, args); err != nil {
			u.Yellow().Println(err)
			os.Exit(1)
		}
	}
}
