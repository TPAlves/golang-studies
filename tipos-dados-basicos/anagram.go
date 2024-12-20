package main

import (
	"fmt"
	"strings"
)

func isAnagrem(word, otherWord string) bool {
	word = strings.ToLower(word)
	otherWord = strings.ToLower(otherWord)
	if len(word) != len(otherWord) {
		return false
	}
	for _, s := range word {
		if !strings.Contains(otherWord, string(s)) {
			return false
		}
	}
	return true
}

func main() {
	word := "amor"
	otherWord := "ramo"
	fmt.Println(isAnagrem(word, otherWord))
}
