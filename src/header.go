package app

import "github.com/charmbracelet/bubbles/textinput"

type header struct {
	inputStyle *Style
	urlInput   textinput.Model
}

func newHeader() header {
	input := textinput.New()
	input.Placeholder = "Enter a URL"
	input.SetValue(initialUrl)
	input.Focus()

	header := header{
		urlInput:   input,
		inputStyle: TextFieldStyle(),
	}
	return header
}
