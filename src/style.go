package app

import "github.com/charmbracelet/lipgloss"

var (
	httpMethodWidth = 10
)

type Style struct {
	BorderColor lipgloss.Color
	InputField  lipgloss.Style
}

var spinnerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("63"))

func UrlInputStyle() *Style {
	s := new(Style)
	s.BorderColor = lipgloss.Color("0")
	s.InputField = lipgloss.NewStyle().
		Foreground(lipgloss.Color("2")).
		BorderStyle(lipgloss.NormalBorder()).
		Padding(0).UnsetMaxWidth()

	return s
}

func MethodInputStyle() *Style {
	s := new(Style)
	s.BorderColor = lipgloss.Color("0")
	s.InputField = lipgloss.NewStyle().
		Foreground(lipgloss.Color("2")).
		BorderStyle(lipgloss.NormalBorder()).
		Padding(0).
		Width(httpMethodWidth)

	return s
}
