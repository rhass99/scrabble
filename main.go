package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readLines(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, strings.TrimSpace(strings.ToLower(scanner.Text())))
	}
	return lines, scanner.Err()
}

func score(word string) int {
	scores := map[string]int{"a": 1, "c": 3, "b": 3, "e": 1, "d": 2, "g": 2,
		"f": 4, "i": 1, "h": 4, "k": 5, "j": 8, "m": 3,
		"l": 1, "o": 1, "n": 1, "q": 10, "p": 3, "s": 1,
		"r": 1, "u": 1, "t": 1, "w": 4, "v": 4, "y": 4,
		"x": 8, "z": 10,
	}
	totalScore := 0
	for _, c := range word {
		ch := string(c)
		totalScore += scores[ch]
	}
	return totalScore
}

func calculateScore(winningWords []string) map[string]int {
	finalScore := map[string]int{}

	for _, word := range winningWords {
		finalScore[word] = score(word)
	}
	return finalScore
}

func validWords(word string, rack string) bool {

	for _, letter := range word {
		if !strings.Contains(rack, string(letter)) {
			return false
		}
		rack = strings.Replace(rack, string(letter), "", 1)
	}
	return true
}

func getWinners(rack string, filename string) []string {
	list, _ := readLines(filename)
	var winningWords []string
	for _, l := range list {
		if validWords(l, rack) {
			winningWords = append(winningWords, l)
		}
	}
	return winningWords
}

func main() {
	var rack string
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./scrabble [RACK]")
		os.Exit(1)
	} else {
		rack = strings.ToLower(os.Args[1])
	}
	winners := getWinners(rack, "sowpods.txt")

	finalScore := calculateScore(winners)

	for k, v := range finalScore {
		fmt.Printf("%s: %d \n", k, v)
	}
}
