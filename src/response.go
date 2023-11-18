package app

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/alecthomas/chroma/quick"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
)

const (
	initial = iota
	loading
	success
	failure
)

type response struct {
	responseBytes []byte
	response      string
	error         error
	status        int
	model         viewport.Model
	width         int
	height        int
}

func (m *response) setViewPortSize(w, h int) {
	m.width = w
	m.height = h
}

func (r *response) responseView() string {
	if r.status == initial {
		return ""
	}

	if r.status == loading {
		return "Loading.."
	}

	if r.status == failure {
		return "Error: " + r.error.Error()
	}

	// Parse the JSON
	var jsonData interface{}
	err := json.Unmarshal(r.responseBytes, &jsonData)
	if err != nil {
		log.Println("Error parsing JSON:", err)
		r.response = "err: " + err.Error()
		return r.response
	}
	// Format the JSON
	formattedJSON, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		log.Println("Error formatting JSON:", err)
	}

	// Syntax highlighting
	var buf bytes.Buffer
	err = quick.Highlight(&buf, string(formattedJSON), "json", "terminal", "monokai")
	if err != nil {
		log.Println("Error highlighting JSON:", err)
	}

	r.response = buf.String()

	return r.response
}

func (r *response) Render() string {
	res := r.responseView()

	focusedModelStyle := lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("0")).
		Width(20).Height(20)

	// r.model.SetContent(focusedModelStyle.Render(res))
	// r.model.Height = 14
	// r.model.Width = r.width

	// return r.model.View()
	return focusedModelStyle.Render(res)
}
