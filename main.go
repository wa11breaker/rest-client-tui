package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	Initialized = iota
	Loading
	Success
	Failed
)

type model struct {
	url      string
	response struct{}
	status   int
}

func newModel() model {
	return model{}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	return "asdf asdf"
}

func main() {
	p := tea.NewProgram(newModel())
	if err := p.Start(); err != nil {
		fmt.Println(err)
		return
	}
}
