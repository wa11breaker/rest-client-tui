package app

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const initialUrl = "http://date.jsontest.com/"

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type OnApiSuccess []byte
type OnApiError struct{ err error }

func (m model) makeApiRequest() tea.Msg {
	return makeRequest(m.header.urlInput.Value())
}

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

func NewModel() model {
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
		m.header.setViewPortSize(msg.Width, msg.Height)
		m.response.setViewPortSize(msg.Width, msg.Height)

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
	items := []list.Item{
		item{title: "Raspberry Pi’s", desc: "I have ’em all over my house"},
		item{title: "Nutella", desc: "It's good on toast"},
		item{title: "Bitter melon", desc: "It cools you down"},
		item{title: "Nice socks", desc: "And by that I mean socks without holes"},
		item{title: "Eight hours of sleep", desc: "I had this once"},
		item{title: "Cats", desc: "Usually"},
	}
	m.list = list.New(items, list.NewDefaultDelegate(), 20, 14)
	m.list.Title = "Request"

	if m.width == 0 {
		return "Loading.."
	}

	m.model.Width = 50
	m.model.Height = m.height
	m.model.SetContent(docStyle.Render(m.list.View()))

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Top,

		lipgloss.JoinHorizontal(
			lipgloss.Left,
			m.model.View(),

			lipgloss.JoinVertical(
				lipgloss.Left,
				m.header.View(),
				// m.response.Render(),
			),
		),
	)
}
