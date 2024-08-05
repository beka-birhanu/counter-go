package main

import (
	"regexp"
	"strings"
)

// Palindrome checks if a given string is a palindrome.
func Palindrome(s string) bool {
	s = strings.ToLower(s)

	re := regexp.MustCompile(`[^\w]`)
	s = re.ReplaceAllString(s, "")

	l := 0
	r := len(s) - 1
	for l <= r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

// WordCount returns a map with the frequency count of each word in the given string.
func WordCount(s string) map[string]int {
	s = strings.ToLower(s)

	re := regexp.MustCompile(`[^\w\s]`)
	s = re.ReplaceAllString(s, "")

	words := strings.Fields(s)
	counts := make(map[string]int)

	for _, word := range words {
		counts[word]++
	}

	return counts
}

