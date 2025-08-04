package cli

import (
	"devsyringe/internal/exceptions"
	process "devsyringe/internal/process"
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	items               []string
	cursor              int
	tickEvery           time.Duration
	pm                  *process.ProcManager
	deleteAllConfirming bool
	showHelp            bool
}

type tickMsg struct{}

func (m model) Init() tea.Cmd {
	return tea.Batch(m.tick())
}

func (m model) tick() tea.Cmd {
	return tea.Tick(m.tickEvery, func(t time.Time) tea.Msg {
		return tickMsg{}
	})
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tickMsg:
		if !m.deleteAllConfirming && !m.showHelp {
			m.items = generateItems(m.pm)
		}
		return m, m.tick()

	case tea.KeyMsg:
		key := msg.String()

		if key == "?" {
			m.showHelp = !m.showHelp
			m.items = generateItems(m.pm)
		}

		if key == "ctrl+c" || key == "q" {
			return m, tea.Quit
		}

		if m.showHelp {
			if key == "esc" {
				m.showHelp = false
			}
			return m, nil
		}

		if m.deleteAllConfirming {
			switch key {
			case "y", "Y":
				m.pm.DeleteAllProcesses()
				m.deleteAllConfirming = false
				m.items = generateItems(m.pm)
			case "n", "N":
				m.deleteAllConfirming = false
				m.items = generateItems(m.pm)
			}
			return m, nil
		}
		switch key {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.items)-1 {
				m.cursor++
			}
		case "S":
			item := m.items[m.cursor]
			m.pm.StopProcess(extractTitleFromItem(item))
			m.items = generateItems(m.pm)
		case "D":
			item := m.items[m.cursor]
			m.pm.DeleteProcess(extractTitleFromItem(item))
		case "alt+D":
			m.deleteAllConfirming = true
		}
	}
	return m, nil
}

func extractTitleFromItem(item string) string {
	return strings.Fields(item)[1]
}

func (m model) View() string {
	s := "\033[H\033[2J"

	if m.deleteAllConfirming {
		s += "Are you sure, you want to delete (and stop) ALL processes? (y/n)"
		return s
	}

	if m.showHelp {
		s += "Help - Keyboard Shortcuts\n\n" +
			"↑ / k\tMove up\n" +
			"↓ / j\tMove down\n" +
			"S\tStop process\n" +
			"D\tDelete process\n" +
			"Alt+D\tDelete ALL processes\n" +
			"?\tToggle help\n" +
			"q\tQuit\n\n" +
			"Press '?' again to return.\n"
		return s
	}

	s += "Processes:\n\n"

	for i, item := range m.items {
		cursor := " "
		if i == m.cursor {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, item)
	}

	s += "\nPress q to quit.\n"
	return s
}

func generateItems(pm *process.ProcManager) []string {
	var items []string
	for _, proc := range pm.GetProcesses() {
		items = append(
			items,
			fmt.Sprintf("%d\t%s\t(%s)", proc.PID, proc.Title, proc.Status),
		)
	}
	return items
}

func tui(pm *process.ProcManager) {
	p := tea.NewProgram(model{
		tickEvery:           500 * time.Millisecond,
		items:               generateItems(pm),
		pm:                  pm,
		deleteAllConfirming: false,
	})
	_, err := p.Run()
	exceptions.Check(err)
}
