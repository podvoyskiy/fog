package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	f "github.com/podvoyskiy/fog/filters"
	"github.com/podvoyskiy/fog/history"
	u "github.com/podvoyskiy/fog/utils"
)

const defaultLimit = 10

type AppConfig struct {
	pathToFile string
	Limit      uint8
	Filter     f.Filtering
}

func Load(configDir string) (*AppConfig, error) {
	configFile := filepath.Join(configDir, "fog", "config")

	if err := os.MkdirAll(filepath.Dir(configFile), 0755); err != nil {
		return nil, err
	}

	config := &AppConfig{
		pathToFile: configFile,
		Limit:      defaultLimit,
		Filter:     f.Default(),
	}

	if _, err := os.Stat(config.pathToFile); os.IsNotExist(err) {
		if err := config.createDefaultConfig(); err != nil {
			return nil, fmt.Errorf("failed to create config file: %s | %w", config.pathToFile, err)
		}
	}

	if err := config.loadFromConfig(); err != nil {
		return nil, err
	}

	return config, nil
}

func (c *AppConfig) Update() error {
	content := fmt.Sprintf("limit=%d\n", c.Limit)

	return os.WriteFile(c.pathToFile, []byte(content), 0644)
}

func (c *AppConfig) PrintStats() error {
	history, err := history.Load()
	if err != nil {
		return err
	}

	u.Cyan().Bold().Println("Most used commands:")

	f := &f.FrequencyFilter{}
	topCommands := f.All(history.Commands)

	for i, cmd := range topCommands {
		if i >= 40 {
			break
		}
		fmt.Printf("%d: %s (%s times)\n", i+1, history.Commands[cmd.Index], u.White().Bold().Sprint(cmd.Score))
	}

	return nil
}

func (c *AppConfig) PrintHelp() {
	u.Yellow().Underline().Println("Options:")
	fmt.Printf("%s            Show this help\n", u.Blue().Sprint("  -h, --help"))
	fmt.Printf("%s           Most used commands\n", u.Blue().Sprint("  -s, --stats"))
	fmt.Printf("%s     Limit results to NUM (current: %d)\n", u.Blue().Sprint("  -l, --limit {NUM}"), c.Limit)
}

func (c *AppConfig) ResetToDefaults() {
	c.Limit = defaultLimit
	c.Filter = f.Default()
}

func (c *AppConfig) createDefaultConfig() error {
	c.ResetToDefaults()
	return c.Update()
}

func (c *AppConfig) loadFromConfig() error {
	content, err := os.ReadFile(c.pathToFile)
	if err != nil {
		return err
	}

	if len(content) == 0 {
		return fmt.Errorf("config file %q is empty", c.pathToFile)
	}

	lines := strings.Split(string(content), "\n")

	for lineNum, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.Split(line, "=")
		if len(parts) != 2 {
			return fmt.Errorf("invalid format (expected key=value). Row: %d", lineNum+1)
		}

		key := strings.TrimSpace(parts[0])
		value, err := u.Uint8(parts[1])
		if err != nil {
			return err
		}

		switch key {
		case "limit":
			c.Limit = value
		default:
			return fmt.Errorf("unknown config key: %s", key)
		}
	}
	return nil
}
