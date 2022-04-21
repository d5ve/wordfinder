package main

import (
	"reflect"
	"strings"
	"testing"
)

var dict string

func TestLoadWords(t *testing.T) {
	words := LoadWords()
	if len(words) < 1000 {
		t.Fatal("Less than 1000 words loaded from dictionary")
	}
	if len(words) > 1000000 {
		t.Fatal("Less than 1000 words loaded from dictionary")
	}
	dict = strings.Join(words, " ")
}

func TestFindWords(t *testing.T) {

	wordsTests := []struct {
		input    string
		expected []string
	}{
		{"a", []string{"a"}},
		{"z", []string{}},
		{"dgo", []string{"do", "dog", "go", "god", "o", "od", "og"}},
	}

	for _, tt := range wordsTests {
		t.Run(tt.input, func(t *testing.T) {
			got := FindWords(tt.input, dict)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("'%s' got %s want %s", tt.input, got, tt.expected)
			}
		})
	}

}
