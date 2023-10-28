package app

import (
	"io/ioutil"
	"net/http"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func makeRequest(url string) tea.Msg {
	c := &http.Client{Timeout: 10 * time.Second}
	res, err := c.Get(url)
	if err != nil {
		return OnApiError{err}
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return OnApiError{err}
	}

	responseString := string(body)
	return OnApiSuccess(responseString)
}
