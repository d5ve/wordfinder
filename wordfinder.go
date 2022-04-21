package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"
)

var saneInputLen = 26

func main() {
	words := loadWords()
	//fmt.Println(words)
	dict := strings.Join(words, " ")
	// fmt.Println(dict)

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	http.HandleFunc("/wordfinder/", wordfinderHandler(dict))
}

func wordfinderHandler(dict string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get a sane-length string of alpha characters.
		chars := r.URL.Path[len("/wordfinder/"):]
		nonAlpha := regexp.MustCompile(`[^a-z]`)
		chars = string(nonAlpha.ReplaceAll([]byte(chars), []byte("")))
		if len(chars) < 1 {
			w.WriteHeader(400)
			return
		}
		if len(chars) > saneInputLen {
			w.WriteHeader(400)
			return
		}

		// Get the list of words made up of the input chars.
		words := findWords(chars, dict)

		// Return the list as JSON.
		jwords, err := json.Marshal(words)
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
			w.WriteHeader(500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jwords)
	}
}

func findWords(input string, dict string) (words []string) {
	return words

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
