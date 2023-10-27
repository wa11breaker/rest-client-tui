package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/alecthomas/chroma/quick"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var spinnerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("63"))

const initialUrl = "http://date.jsontest.com/"

func makeRequest(url string) tea.Msg {
	c := &http.Client{Timeout: 10 * time.Second}
	res, err := c.Get(url)
	if err != nil {
		return OnApiError{err}
	}
	defer res.Body.Close() // nolint:errcheck

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return OnApiError{err}
	}

	responseString := string(body)
	return OnApiSuccess(responseString)
}

type OnApiSuccess []byte
type OnApiError struct{ err error }

func (m model) makeApiRequest() tea.Msg {
	return makeRequest(m.urlInput.Value())
}

type Style struct {
	BorderColor lipgloss.Color
	InputField  lipgloss.Style
}

func DefaultStyle() *Style {
	s := new(Style)
	s.BorderColor = lipgloss.Color("0")
	s.InputField = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).BorderStyle(lipgloss.NormalBorder()).Padding(1).Width(50)
	return s
}

type model struct {
	urlInput   textinput.Model
	responce   string
	inputStyle *Style
	width      int
	height     int
}

func NewModel() model {
	input := textinput.New()
	input.Placeholder = "Enter a URL"
	input.SetValue(initialUrl)
	input.Focus()

	s := spinner.New()
	s.Style = spinnerStyle
	model := model{
		urlInput:   input,
		inputStyle: DefaultStyle(),
	}
	return model
}

func (m model) Init() tea.Cmd {
	return nil
}

type customCmd func() tea.Msg

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
			m.responce = "err: " + err.Error()
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

		m.responce = buf.String()

	case OnApiError:
		m.responce = "err: " + string(msg.err.Error())
	}
	m.urlInput, _ = m.urlInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.width == 0 {
		return "Loading.."
	}

	// string := []string{
	// 	m.inputStyle.InputField.Render(m.urlInput.View()),
	// }

	// if len(m.responce) > 0 {
	// 	string = append(string, m.responce)
	// }

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,

		lipgloss.JoinVertical(
			lipgloss.Center,
			m.inputStyle.InputField.Render(m.urlInput.View()),
			m.responce,
		),
	)
}

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatal("err: %w", err)
	}
	defer f.Close()

	p := tea.NewProgram(NewModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
