package app

const (
	initial = iota
	loading
	success
	failure
)

type response struct {
	response string
	error    error
	status   int
}
