package main

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/iskaa02/qalam"
)

func typerUpdate(m *model, msg tea.Msg) (model, tea.Cmd) {
	if m.Start.IsZero() {
		m.Start = time.Now()
	}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {

		case tea.KeyCtrlC:
			return *m, tea.Quit
		// ctrl+h maps to ctrl+backspace
		// use ctrl+backspace to delete whole word
		case tea.KeyCtrlH:
			m.CurrentWord = ""
		case tea.KeyBackspace:
			if len(m.CurrentWord) >= 1 {
				m.CurrentWord = m.CurrentWord[:len(m.CurrentWord)-1]
			}
		case tea.KeyEscape:
			setEnded(m)
			break
		case tea.KeyRunes:
			if len(msg.String()) > 1 {
				break
			}
			if msg.String() == " " {
				handleSpace(m)
				break
			}
			m.CurrentWord += msg.String()
		}
	}
	return *m, nil
}

func typerView(m model) string {
	styler := qalam.NewStyler()
	s := styler.Dim().Sprintf("> %s", m.CurrentWord) + "\n\n"
	styler.Remove("gray")

	for i, word := range m.Words {
		if i == m.Index {
			s += styleCurrentWord(word, m.CurrentWord) + " "
			continue
		}
		if m.Index > i {
			isCorrect := m.PreviousWords[i]
			if isCorrect {
				s += styler.Green().Sprintf("%s", word) + " "
				styler.Remove("green")
				continue
			}
			s += styler.Red().Sprintf("%s", word) + " "
			styler.Remove("red")

			continue
		}
		s += styler.Sprintf("%s", word) + " "
	}
	s = wrapWords(s)
	return s
}

func handleSpace(m *model) {
	if len(m.CurrentWord) < 1 {
		return
	}
	m.PreviousWords[m.Index] = m.Words[m.Index] == m.CurrentWord
	if m.Index == len(m.Words)-1 {
		setEnded(m)
	} else {
		m.Index++
	}
	m.CurrentWord = ""
	return
}

func setEnded(m *model) {
	m.End = time.Now()
	m.Ended = true
}

func styleCurrentWord(actual, typed string) string {
	styler := qalam.NewStyler().Blue().Bold()
	s := ""
	ok := matchString(actual, typed)
	if ok {
		s += styler.Green().Sprintf("%s", actual[:len(typed)])
		styler.Remove("green")
		s += styler.Sprintf("%s", actual[len(typed):])
	} else {
		styler.Red()
		if len(actual) < len(typed) {
			s += styler.Sprintf("%s", actual)
		} else {
			s += styler.Sprintf("%s", actual[:len(typed)])
			styler.Remove("red")
			s += styler.Sprintf("%s", actual[len(typed):])
		}
	}
	return s
}
