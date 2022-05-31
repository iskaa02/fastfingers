package main

import (
	"flag"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	wordsLength      int
	lang             string
	resourceFileName string
)

func main() {
	flag.IntVar(&wordsLength, "w", 50, "Number of words in test.")
	flag.StringVar(&lang, "l", "english", "Predefined test 'english, english100, english1000, english2000'")
	flag.StringVar(&resourceFileName, "f", "", "Resource file to use, words should be seperated by new lines.")
	flag.Parse()
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
