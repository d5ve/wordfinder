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
	words := LoadWords()
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
		words := FindWords(chars, dict)

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

func FindWords(chars string, dict string) (words []string) {

	// Track the frequency of each char in the input.
	freq := make(map[rune]int)
	for _, c := range chars {
		f, exists := freq[c]
		if exists {
			freq[c] = f + 1
		} else {
			freq[c] = 1

		}
	}

	// Find all words in the dictionary that can be made solely from the
	// input chars. Also filter to matches no longer than the input chars.
	// [di] shouldn't match did
	var re = regexp.MustCompile(fmt.Sprintf(`\b([%s]{1,%d})\b`, chars, len(chars)))
	matches := re.FindAllString(dict, -1)

	// Process each match further to check that the character frequency is
	// no larger than that of the input.
match:
	for _, word := range matches {
		wfreq := make(map[rune]int)
		for _, c := range word {
			wf, exists := wfreq[c]
			if exists {
				f, exists := freq[c]
				if exists && wf+1 > f {
					continue match
				}
				wfreq[c] = wf + 1
			} else {
				wfreq[c] = 1

			}
		}
		words = append(words, word)
	}
	return words

}

func LoadWords() []string {
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
