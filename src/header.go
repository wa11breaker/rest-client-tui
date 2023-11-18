package app

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

type header struct {
	inputStyle      *Style
	methoInputStyle *Style
	urlInput        textinput.Model
	methodInput     textinput.Model
	width           int
	height          int
}

func (m *header) setViewPortSize(w, h int) {
	m.width = w
	m.height = h

	m.urlInput.Width = m.width - 50 - (httpMethodWidth + 7)
}

func newHeader() header {
	input := textinput.New()
	input.Placeholder = "Enter a URL"
	input.SetValue(initialUrl)
	// input.Focus()

	methodInput := textinput.New()
	methodInput.Placeholder = "GET"
	methodInput.SetValue("GET")

	header := header{
		methodInput:     methodInput,
		urlInput:        input,
		inputStyle:      UrlInputStyle(),
		methoInputStyle: MethodInputStyle(),
	}
	return header
}

func (m header) View() string {
	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		m.methoInputStyle.InputField.Render(m.methodInput.View()),
		m.inputStyle.InputField.Render(m.urlInput.View()),
	)
}
