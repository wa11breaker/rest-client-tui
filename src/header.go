package app

import "github.com/charmbracelet/bubbles/textinput"

type header struct {
	inputStyle      *Style
	methoInputStyle *Style
	urlInput        textinput.Model
	methodInput     textinput.Model
}

func newHeader() header {
	input := textinput.New()
	input.Placeholder = "Enter a URL"
	input.SetValue(initialUrl)
	input.Focus()

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
