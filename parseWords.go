package main

import (
	_ "embed"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/iskaa02/qalam/bbcode"
)

//go:embed resources/english.txt
var english []byte

//go:embed resources/english1000.txt
var english1000 []byte

//go:embed resources/english2000.txt
var english2000 []byte

func parseWords() []string {
	f := english
	switch lang {
	case "english":
	case "english1000":
		f = english1000
	case "english200":
		f = english2000
	default:
		bbcode.Printf("[red]-l flag[/red] should only be one of those values [b red]english english1000 english2000[/b red]\n")
		os.Exit(1)
	}
	if resourceFileName != "" {
		file, err := os.ReadFile(resourceFileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		f = file
	}
	words := strings.Split(string(f), "\n")
	return randomPick(wordsLength, words)
}

func randomPick(length int, words []string) []string {
	if length > len(words) {
		panic("pick lengh more than words")
	}
	newWords := []string{}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(words), func(i, j int) { words[i], words[j] = words[j], words[i] })
	for i := 0; i < length; i++ {
		index := rand.Intn(len(words))
		if words[index] == "" {
			continue
		}
		newWords = append(newWords, words[index])
	}
	return newWords
}
