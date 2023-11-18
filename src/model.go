package app

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
)

type model struct {
	list     list.Model
	model    viewport.Model
	response response
	header   header
	width    int
	height   int
}
