package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	words := loadWords()
	//fmt.Println(words)
	str := strings.Join(words, " ")
	fmt.Println(str)

}

func loadWords() []string {
	bytes, err := ioutil.ReadFile("/usr/share/dict/words")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	str := strings.ToLower(string(bytes))
	lines := strings.Split(str, "\n")
	var words []string

	var atoz = regexp.MustCompile(`^[a-z]+$`)
	var single = regexp.MustCompile(`[aio]`)

	// Filter out words that contain characters outside a-z. Plus any
	// erroneous single-char words.
	filter := func(word string) bool {
		if !atoz.MatchString(word) {
			// Filter out words with non-a-z.
			return false
		}
		if len(word) == 1 {
			if single.MatchString(word) {
				return true
			}
			return false
		}
		return true
	}
	seen := make(map[string]bool) // Used to filter duplicates.
	for _, word := range lines {
		if _, s := seen[word]; !s {
			seen[word] = true
			if filter(word) {
				words = append(words, word)
			}
		}
	}

	sort.Strings(words)
	return words
}
