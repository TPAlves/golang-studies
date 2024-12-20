package main

import (
	"fmt"
	"testing"
)

func TestAnagram(t *testing.T) {
	testAnagrams := []struct {
		word      string
		otherWord string
		result    bool
	}{
		{word: "amor", otherWord: "ramo", result: true},
		{word: "listen", otherWord: "silent", result: true},
		{word: "evil", otherWord: "vile", result: true},
		{word: "fluster", otherWord: "restful", result: true},
		{word: "god", otherWord: "dog", result: true},
		{word: "evil", otherWord: "vile", result: true},
		{word: "fluster", otherWord: "restful", result: true},
		{word: "uva", otherWord: "vuma", result: false},
		{word: "aabb", otherWord: "ab", result: false},
	}
	for _, tt := range testAnagrams {
		t.Run(fmt.Sprintf("1º Palavra: %s, 2º Palavra: %s", tt.word, tt.otherWord), func(t *testing.T) {
			result := isAnagrem(tt.word, tt.otherWord)
			if result != tt.result {
				t.Errorf("%s - %s, result: %t, expered: %t", tt.word, tt.otherWord, result, tt.result)
			}
		})
	}
}
