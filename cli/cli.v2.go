package cli

import (
	"mesh/core"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	Home = iota
	Computes
	Config
)

type Model struct {
	List       []string
	ListCursor int

	Command string
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) View() string {
	
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl + c":
			return m, tea.Quit
		case "enter":
			m.Command = ""
		default:
			m.Command += msg.String()
			return m, nil
		}
	}
}

func Init(m tea.Model) error {
	p := tea.NewProgram(m)
	_, err := p.Run()
	core.Logln("[E] Cli init ended with error.\n[EMSG]: ", err)
	return err
}
