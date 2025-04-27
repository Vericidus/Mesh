package main

import (
	"fmt"
	"mesh/core"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	tasks     []string
	selected  int
	view      string
	statusMsg string
}

func initialModel() model {
	return model{
		tasks:     []string{"extract_data", "transform", "load_data"},
		selected:  0,
		view:      "list",
		statusMsg: "Running...",
	}
}

type tickMsg time.Time

func (m model) Init() tea.Cmd {
	return tea.Batch(tick(), nil)
}

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up":
			if m.selected > 0 {
				m.selected--
			}
		case "down":
			if m.selected < len(m.tasks)-1 {
				m.selected++
			}
		case "enter":
			m.view = "detail"
		case "esc":
			m.view = "list"
		}

	case tickMsg:
		m.statusMsg = fmt.Sprintf("Updated at %s", time.Now().Format("15:04:05"))
		return m, tick()
	}

	return m, nil
}

func (m model) View() string {
	if m.view == "detail" {
		return fmt.Sprintf(
			"Task: %s\n\nDetails coming soon...\n\n[esc to go back]",
			m.tasks[m.selected],
		)
	}

	s := "DAG: my-dag\n\n"
	for i, task := range m.tasks {
		cursor := "  "
		if i == m.selected {
			cursor = "> "
		}
		s += fmt.Sprintf("%s%s\n", cursor, task)
	}
	s += fmt.Sprintf("\n%s\n", m.statusMsg)
	s += "\n[↑/↓ to move, enter for details, q to quit]"
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		core.Logln("error:", err)
		os.Exit(1)
	}
}
