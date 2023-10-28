package app

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/alecthomas/chroma/quick"
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
	response string
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
			return m, m.makeApiRequest
		}

	case OnApiSuccess:
		var jsonData interface{}
		err := json.Unmarshal(msg, &jsonData)
		if err != nil {
			log.Println("Error parsing JSON:", err)
			m.response = "err: " + err.Error()
			return m, nil
		}

		formattedJSON, err := json.MarshalIndent(jsonData, "", "  ")
		if err != nil {
			log.Println("Error formatting JSON:", err)
		}

		var buf bytes.Buffer

		// Syntax highlighting using chroma
		err = quick.Highlight(&buf, string(formattedJSON), "json", "terminal", "monokai")
		if err != nil {
			log.Println("Error highlighting JSON:", err)
		}

		m.response = buf.String()

	case OnApiError:
		m.response = "err: " + string(msg.err.Error())
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
			lipgloss.Center,
			m.header.inputStyle.InputField.Render(m.header.urlInput.View()),
			m.response,
		),
	)
}
