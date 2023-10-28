package main

import (
	"log"
	"rest-client-tui/src"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatal("err: %w", err)
	}
	defer f.Close()

	p := tea.NewProgram(app.NewModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
