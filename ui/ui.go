package ui

import (
	"fmt"
	"os"

	c "github.com/podvoyskiy/fog/config"
	s "github.com/podvoyskiy/fog/searcher"

	tea "github.com/charmbracelet/bubbletea"
	u "github.com/podvoyskiy/fog/utils"
)

type model struct {
	searcher *s.Searcher
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	key, ok := msg.(tea.KeyMsg)
	if !ok {
		return m, nil
	}

	switch key.String() {

	case "ctrl+c", "esc":
		return m, tea.Quit

	case "enter":
		if cmd, ok := m.searcher.GetSelectedCommand(); ok {
			fmt.Printf("history -s \"%s\"\n%s", cmd, cmd)
			return m, tea.Quit
		}

	case "right":
		if cmd, ok := m.searcher.GetSelectedCommand(); ok {
			fmt.Printf("echo '%s'", cmd)
			return m, tea.Quit
		}

	case "up":
		if m.searcher.SelectedIndex > 0 {
			m.searcher.SelectedIndex--
		}

	case "down":
		if m.searcher.SelectedIndex < m.searcher.ResultCount()-1 {
			m.searcher.SelectedIndex++
		}

	case "backspace":
		if len(m.searcher.SearchQuery) > 0 {
			m.searcher.SearchQuery = m.searcher.SearchQuery[:len(m.searcher.SearchQuery)-1]
			m.searcher.ApplyFilter()
		}

	default:
		if len(key.String()) == 1 {
			m.searcher.SearchQuery += key.String()
			m.searcher.ApplyFilter()
		}
	}

	return m, nil
}

func (m model) View() string {
	prompt := u.Cyan().Sprint("> " + m.searcher.SearchQuery)
	selected := u.Green().Bold().Sprint

	out := prompt + "\n"

	count := m.searcher.ResultCount()

	for i := range count {
		if cmd, ok := m.searcher.GetCommandByIndex(i); ok {
			if i == m.searcher.SelectedIndex {
				out += "> " + selected(cmd) + "\n"
			} else {
				out += "  " + cmd + "\n"
			}
		}
	}

	return out
}

func Run(config *c.AppConfig) error {
	s, err := s.Init(config.Filter, config.Limit)
	if err != nil {
		return err
	}

	prog := tea.NewProgram(model{searcher: s},
		tea.WithAltScreen(),
		tea.WithOutput(os.Stderr), // TUI goes to stderr so stdout only contains the selected command for eval "$(cmd)"
	)

	_, err = prog.Run()
	return err
}
