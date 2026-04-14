package cmds

import (
	"fmt"
	"strings"

	c "github.com/podvoyskiy/fog/config"
	u "github.com/podvoyskiy/fog/utils"
)

func HandleCmd(config *c.AppConfig, args []string) error {
	cmd := strings.ToLower(args[0])

	switch cmd {
	case "--limit", "-l":
		if len(args) < 2 {
			return fmt.Errorf("missing value for %q", cmd)
		}

		value, err := u.Uint8(args[1])
		if err != nil {
			return err
		}
		if value == 0 {
			return fmt.Errorf("limit cannot be 0")
		}
		config.Limit = value

		if err := config.Update(); err != nil {
			return err
		}
		u.Green().Printf("updated settings %s | value: %d\n", cmd, value)
		return nil
	case "help", "-h", "--help":
		config.PrintHelp()
		return nil
	default:
		return fmt.Errorf("unknown command: %s. use --help", cmd)
	}
}
