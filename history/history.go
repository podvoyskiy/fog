package history

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

type CommandHistory struct {
	Commands []string
}

func Load() (*CommandHistory, error) {
	usr, err := user.Current()
	if err != nil {
		return nil, err
	}

	historyPath := filepath.Join(usr.HomeDir, ".bash_history")
	info, err := os.Stat(historyPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("file ~/.bash_history not found")
		}
		return nil, err
	}

	if info.Size() == 0 {
		return nil, fmt.Errorf("file ~/.bash_history is empty")
	}

	file, err := os.Open(historyPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			lines = append(lines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	commands := make([]string, 0, len(lines))

	for i := len(lines) - 1; i >= 0; i-- {
		commands = append(commands, lines[i])
	}

	return &CommandHistory{Commands: commands}, nil
}
