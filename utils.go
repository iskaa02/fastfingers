package main

import (
	"math"
	"os"

	"github.com/iskaa02/qalam/gradient"
	"github.com/muesli/reflow/indent"
	"github.com/muesli/reflow/padding"
	"github.com/muesli/reflow/wordwrap"
	"golang.org/x/term"
)

func generateWords() []string {
	return parse()
}

func wrapWords(s string) string {
	w, _, err := term.GetSize(int(os.Stdout.Fd()))
	w = int(math.Min(float64(w), 100))
	if err != nil {
		w = 60
	}
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
	s := gradient.Vice("b", "i", "u").Apply("Fast Fingers") + "\n\n"
	s += body
	s = padding.String(s, 3)
	s = indent.String(s, 1)
	return s
}
