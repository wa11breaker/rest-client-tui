package app

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const initialUrl = "http://date.jsontest.com/"

type OnApiSuccess []byte
type OnApiError struct{ err error }

func (m model) makeApiRequest() tea.Msg {
	return makeRequest(m.header.urlInput.Value())
}

type model struct {
	response response
	header   header
	width    int
	height   int
}

func NewModel() model {

	s := spinner.New()
	s.Style = spinnerStyle
	header := newHeader()
	model := model{
		header: header,
	}
	return model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			m.response.status = loading
			return m, m.makeApiRequest
		}

	case OnApiSuccess:
		m.response.status = success
		m.response.responseBytes = msg

	case OnApiError:
		m.response.status = failure
		m.response.error = msg.err

	}
	m.header.urlInput, _ = m.header.urlInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.width == 0 {
		return "Loading.."
	}

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,

		lipgloss.JoinVertical(
			lipgloss.Left,
			lipgloss.JoinHorizontal(
				lipgloss.Left,
				m.header.methoInputStyle.InputField.Render(m.header.methodInput.View()),
				m.header.inputStyle.InputField.Render(m.header.urlInput.View()),
			),
			m.response.render(),
		),
	)
}
