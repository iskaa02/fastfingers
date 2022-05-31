package main

import (
	"math"
	"os"
	"strings"

	"github.com/muesli/reflow/indent"
	"github.com/muesli/reflow/padding"
	"github.com/muesli/reflow/wordwrap"
	"golang.org/x/term"
)

const (
	MIN_HEIGHT = 10
	MIN_WIDHT  = 100
)

func generateWords() []string {
	return parse()
}

func getAppHeight() int {
	_, h, err := term.GetSize(int(os.Stdout.Fd()))
	h = int(math.Min(float64(h), float64(MIN_HEIGHT)))
	if err != nil {
		h = MIN_HEIGHT
	}
	return h
}

func getAppWidth() int {
	w, _, err := term.GetSize(int(os.Stdout.Fd()))
	w = int(math.Min(float64(w), MIN_WIDHT))
	if err != nil {
		w = MIN_WIDHT
	}
	return w
}

func wrapWords(s string) string {
	w := getAppWidth()
	f := wordwrap.NewWriter(w)
	f.Write([]byte(s))
	return f.String()
}

// Match the word to what the used typed
func matchString(actual, typed string) bool {
	if len(actual) < len(typed) {
		return false
	}
	for i := 0; i < len(typed); i++ {
		if actual[i] != typed[i] {
			return false
		}
	}
	return true
}

func Layout(body string) string {
	body = CenterHorizontally(body)
	body = CenterVertically(body)
	return body
}

func CenterHorizontally(body string) string {
	termW, _, _ := term.GetSize(int(os.Stdout.Fd()))
	s := body
	appW := getAppWidth()
	s = padding.String(s, uint(termW-appW)/2)
	s = indent.String(s, uint(termW-appW)/2)
	return s
}

func CenterVertically(body string) string {
	_, termH, _ := term.GetSize(int(os.Stdout.Fd()))
	appH := len(strings.Split(body, "\n"))
	diff := termH/2 - appH
	if diff < 0 {
		diff = 1
	}
	body = strings.Repeat("\n", diff) + body
	return body
}
