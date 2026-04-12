package cmds

import (
	"fmt"
	"strings"

	c "github.com/podvoyskiy/fog/config"
	f "github.com/podvoyskiy/fog/filters"
	u "github.com/podvoyskiy/fog/utils"
)

func HandleCmd(config *c.AppConfig, args []string) error {
	cmd := strings.ToLower(args[0])

	switch cmd {
	case "--max_results", "-m", "--filter", "-f":
		if len(args) < 2 {
			return fmt.Errorf("missing value for %q", cmd)
		}

		value, err := u.Uint8(args[1])
		if err != nil {
			return err
		}

		switch cmd {
		case "--max_results", "-m":
			if value == 0 {
				return fmt.Errorf("max_results cannot be 0")
			}
			config.MaxResults = value
		case "--filter", "-f":
			filter, err := f.FromUint8(value)
			if err != nil {
				return err
			}
			config.Filter = filter
		}

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
