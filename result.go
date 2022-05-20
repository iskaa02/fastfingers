package main

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/iskaa02/qalam"
)

type keyMap struct {
	restart key.Binding
	quit    key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.restart, k.quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.restart},
		{k.quit},
	}
}

func resultView(m model) (s string) {
	wpm, correctWords, wrongWords := getScore(m)
	var accuracy float64
	accuracy = float64(correctWords) / float64(wrongWords+correctWords)
	accuracy *= 100.0
	styler := qalam.NewStyler().Bold().Italic()
	s = styler.Sprintf("WPM: %.2f\n", wpm)
	s += styler.Sprintf("Correct Words: %v\n", correctWords)
	s += styler.Sprintf("Wrong words: %v\n\n", wrongWords)
	s += styler.Sprintf("Accuracy: %.2f%%\n\n", accuracy)
	s += m.Help.View(keys)
	return
}

func resultUpdate(m *model, msg tea.Msg) (model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// If we set a width on the help menu it can it can gracefully truncate
		// its view as needed.
		m.Help.Width = msg.Width
	case tea.KeyMsg:

		switch {
		case key.Matches(msg, keys.quit):
			return *m, tea.Quit
		case key.Matches(msg, keys.restart):
			return initialModel(), nil
		}
	}
	return *m, nil
}

var keys = keyMap{
	restart: key.NewBinding(key.WithKeys("r"),
		key.WithHelp("r", "Restart"),
	),
	quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}

func getScore(m model) (wpm float64, correctWords, wrongWords int) {
	correctChars := 0
	for k, v := range m.PreviousWords {
		if v {
			correctChars += len(m.Words[k])
			correctWords++
		} else {
			wrongWords++
		}
	}
	wpm = (float64(correctChars) / charsPerWord) / (float64(m.End.Sub(m.Start).Minutes()))
	return
}
