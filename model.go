package main

import (
	"time"

	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

const charsPerWord = 5

type model struct {
	Words         []string
	PreviousWords map[int]bool
	Index         int
	Help          help.Model
	CurrentWord   string
	Start         time.Time
	End           time.Time
	Ended         bool
	GradientValue float64
}

func initialModel() model {
	words := generateWords()
	return model{
		Words:         words,
		PreviousWords: make(map[int]bool),
		Help:          help.New(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() (s string) {
	if m.Ended {
		s = resultView(m)
	} else {
		s = typerView(m)
	}
	return Layout(s)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.Ended {
		return resultUpdate(&m, msg)
	}
	return typerUpdate(&m, msg)
}
