package main

const (
	listView uint = iota
	titleView
	bodyView
)

type model struct {
	state unit
}

func NewModel() model {
	return model {
		state: listView
	}
}

func (m model) Init ()