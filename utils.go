package main

import (
	"math"
	"os"
	"strings"

	"github.com/iskaa02/qalam/gradient"
	"github.com/muesli/reflow/indent"
	"github.com/muesli/reflow/padding"
	"github.com/muesli/reflow/wordwrap"
	"golang.org/x/term"
)

func generateWords() []string {
	return parse()
}

func getAppWidth() int {
	w, _, err := term.GetSize(int(os.Stdout.Fd()))
	w = int(math.Min(float64(w), 100))
	if err != nil {
		w = 60
	}
	return w
}

func uiSeperator() string {
	grad, _ := gradient.NewGradientBuilder().HtmlColors("#C0FFEE", "#BAD", "#DECAFF").Domain(0.1, 0.3, 0.8).Build()
	return grad.Apply(strings.Repeat("─", getAppWidth())) + "\n"
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
	s := uiSeperator()
	s += body
	s = padding.String(s, 3)
	s = indent.String(s, 1)
	return s
}
