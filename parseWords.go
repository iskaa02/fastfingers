package main

import (
	"math/rand"
	"os"
	"strings"
	"time"
)

func parse() []string {
	file, err := os.ReadFile("english.txt")
	if err != nil {
		panic("Resource file not found")
	}
	words := strings.Split(string(file), "\n")
	return randomPick(50, words)
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
