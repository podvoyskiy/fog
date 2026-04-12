package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	f "github.com/podvoyskiy/fog/filters"
	u "github.com/podvoyskiy/fog/utils"
)

const defaultMaxResults = 10

type AppConfig struct {
	pathToFile string
	MaxResults uint8
	Filter     f.Filter
}

func Load(configDir string) (*AppConfig, error) {
	configFile := filepath.Join(configDir, "fog", "config")

	if err := os.MkdirAll(filepath.Dir(configFile), 0755); err != nil {
		return nil, err
	}

	config := &AppConfig{
		pathToFile: configFile,
		MaxResults: defaultMaxResults,
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
	content := fmt.Sprintf("max_results=%d\nfilter_id=%d\n", c.MaxResults, c.Filter.GetId())

	return os.WriteFile(c.pathToFile, []byte(content), 0644)
}

func (c *AppConfig) PrintHelp() {
	u.Yellow().Underline().Println("Options:")
	fmt.Printf("%s                  Show this help\n", u.Blue().Sprint("  -h, --help"))
	fmt.Printf("%s     Set maximum number of results to display (current: %d)\n", u.Blue().Sprint("  -m, --max_results {NUM}"), c.MaxResults)
	fmt.Printf("%s     Set filter algorithm [1 - SkimFilter, 2 - SubstringFilter] (current: %d)\n",
		u.Blue().Sprint("  -f, --filter      {NUM}"), c.Filter.GetId())
}

func (c *AppConfig) ResetToDefaults() {
	c.MaxResults = defaultMaxResults
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
		case "max_results":
			c.MaxResults = value

		case "filter_id":
			filter, err := f.FromUint8(value)
			if err != nil {
				return err
			}
			c.Filter = filter
		default:
			return fmt.Errorf("unknown config key: %s", key)
		}
	}
	return nil
}
