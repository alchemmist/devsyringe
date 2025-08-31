package tui

import (
	"fmt"
	"strings"
	"time"

	"github.com/alchemmist/devsyringe/internal/exceptions"
	process "github.com/alchemmist/devsyringe/internal/process"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type model struct {
	items               []string
	cursor              int
	tickEvery           time.Duration
	pm                  *process.ProcManager
	showHelp            bool
	table               table.Model
	help                help.Model
	keys                keyMap
	input               textinput.Model
	confirmingDeleteAll bool
	inputSubmitted      bool
	showLogs            bool
	viewport            viewport.Model
	filename            string
	content             string
	ready               bool
}

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Padding(0, 1).
			Border(lipgloss.RoundedBorder()).
			BorderRight(true)

	infoStyle = lipgloss.NewStyle().
			Padding(0, 1).
			Border(lipgloss.RoundedBorder()).
			BorderLeft(true)
)

func (m model) headerView() string {
	title := titleStyle.Render(m.filename)
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m model) footerView() string {
	info := infoStyle.Render(fmt.Sprintf("%3.f%%", m.viewport.ScrollPercent()*100))
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(info)))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type keyMap struct {
	Up        key.Binding
	Down      key.Binding
	Logs      key.Binding
	Stop      key.Binding
	Delete    key.Binding
	DeleteAll key.Binding
	Help      key.Binding
	Quit      key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Up, k.Down, k.Stop, k.Delete, k.DeleteAll, k.Help, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down},
		{k.Stop, k.Delete, k.DeleteAll},
		{k.Help, k.Quit},
	}
}

var keys = keyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "Previous process"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "Next process"),
	),
	Logs: key.NewBinding(
		key.WithKeys("L"),
		key.WithHelp("L", "show logs"),
	),
	Stop: key.NewBinding(
		key.WithKeys("S"),
		key.WithHelp("S", "Stop process"),
	),
	Delete: key.NewBinding(
		key.WithKeys("D"),
		key.WithHelp("D", "Delete process"),
	),
	DeleteAll: key.NewBinding(
		key.WithKeys("alt+D"),
		key.WithHelp("alt+D", "Delete ALL processes"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
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
	case tea.WindowSizeMsg:
		headerHeight := lipgloss.Height(m.headerView())
		footerHeight := lipgloss.Height(m.footerView())
		verticalMarginHeight := headerHeight + footerHeight

		if !m.ready {
			m.viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight)
			m.viewport.YPosition = headerHeight
			m.viewport.SetContent(m.content)
			m.ready = true
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height - verticalMarginHeight
		}

	case tickMsg:
		if !m.inputSubmitted && !m.showHelp {
			m.table.SetRows(generateRows(m.pm))
		}
		return m, m.tick()

	case tea.KeyMsg:
		key := msg.String()

		if key == "l" && !m.showHelp && !m.confirmingDeleteAll {
			row := m.table.SelectedRow()
			if len(row) != 0 {
				title := row[1]
				proc, err := m.pm.GetProcess(title)
				exceptions.Check(err)

				tmp := strings.Split(proc.LogFile, "/")
				m.filename = tmp[len(tmp)-1]
				m.content = proc.GetLogs()
				m.viewport.SetContent(proc.GetLogs())
				m.showLogs = true
				return m, nil
			}
		}

		if m.showLogs {
			switch key {
			case "esc":
				m.showLogs = false
				return m, nil
			case "q", "ctrl+c":
				return m, tea.Quit
			}

			var cmd tea.Cmd
			m.viewport, cmd = m.viewport.Update(msg)
			return m, cmd
		}

		if key == "?" {
			m.showHelp = !m.showHelp
			return m, nil
		}

		if key == "ctrl+c" || key == "q" {
			return m, tea.Quit
		}

		if m.showHelp {
			if key == "esc" || key == "?" {
				m.showHelp = false
			}
			return m, nil
		}

		if m.confirmingDeleteAll {
			switch key {
			case "enter":
				value := m.input.Value()
				if value == "y" || value == "yes" {
					m.pm.DeleteAllProcesses()
					m.table.SetRows(generateRows(m.pm))
				}
				m.input.Reset()
				m.confirmingDeleteAll = false
				return m, nil
			case "esc":
				m.confirmingDeleteAll = false
				m.input.Reset()
				return m, nil
			}

			var cmd tea.Cmd
			m.input, cmd = m.input.Update(msg)
			return m, cmd
		}

		switch key {
		case "S":
			row := m.table.SelectedRow()
			if len(row) != 0 {
				m.pm.StopProcess(row[1])
				m.table.SetRows(generateRows(m.pm))
			}
		case "D":
			row := m.table.SelectedRow()
			if len(row) != 0 {
				m.pm.DeleteProcess(row[1])
				m.table.SetRows(generateRows(m.pm))
			}
		case "alt+D":
			m.confirmingDeleteAll = true
			m.input.Focus()
			m.input.Reset()
			return m, nil
		}

		var cmd tea.Cmd
		m.table, cmd = m.table.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m model) View() string {
	switch {
	case m.showHelp:
		return baseStyle.Render(m.help.View(m.keys)) + "\n"
	case !m.ready:
		return "\n  Initializing..."
	case m.showLogs:
		return fmt.Sprintf("%s\n%s\n%s", m.headerView(), m.viewport.View(), m.footerView())
	case m.confirmingDeleteAll:
		return baseStyle.Render("Are you sure to delete (and stop) ALL processes?\nThe logs will be cleared.\n\n" + m.input.View() + "\n(Press enter to confirm or esc to cancel)")
	default:
		return baseStyle.Render(m.table.View()) + "\n"
	}
}

func generateRows(pm *process.ProcManager) []table.Row {
	rows := []table.Row{}
	for _, proc := range pm.GetProcesses() {
		rows = append(rows, table.Row{
			fmt.Sprintf("%d", proc.PID),
			proc.Title,
			proc.Status.String(),
			proc.Command,
		})
	}
	return rows
}

func generateItems(pm *process.ProcManager) []string {
	var items []string
	for _, proc := range pm.GetProcesses() {
		items = append(
			items,
			fmt.Sprintf("%s (%d)", proc.Title, proc.PID),
		)
	}
	return items
}

func Tui(pm *process.ProcManager) {
	columns := []table.Column{
		{Title: "PID", Width: 8},
		{Title: "Title", Width: 15},
		{Title: "Status", Width: 6},
		{Title: "Run with", Width: 25},
	}

	rows := generateRows(pm)

	maxHeight := 11
	minHeight := 3

	calculatedHeight := len(rows) + 1
	if calculatedHeight > maxHeight {
		calculatedHeight = maxHeight
	} else if calculatedHeight < minHeight {
		calculatedHeight = minHeight
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(calculatedHeight),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("#000")).
		Background(lipgloss.Color("#547d4e")).
		Bold(false)
	t.SetStyles(s)

	width, height := tea.WindowSizeMsg{Width: 100, Height: 30}.Width, tea.WindowSizeMsg{Height: 30}.Height
	vp := viewport.New(width, height-4)
	vp.Style = baseStyle

	helpModel := help.New()
	helpModel.ShowAll = true

	input := textinput.New()
	input.Placeholder = "Type yes (y) to delete all"
	input.Focus()
	input.CharLimit = 20
	input.Width = 30

	p := tea.NewProgram(model{
		tickEvery: 500 * time.Millisecond,
		items:     generateItems(pm),
		pm:        pm,
		table:     t,
		help:      helpModel,
		keys:      keys,
		input:     input,
		viewport:  vp,
	}, tea.WithMouseCellMotion())
	fmt.Print("\0337")
	_, err := p.Run()
	exceptions.Check(err)
	fmt.Print("\0338")
	fmt.Print("\033[J")
}
