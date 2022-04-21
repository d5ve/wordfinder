package main

import "testing"

var words []string

func TestLoadWords(t *testing.T) {
	words = LoadWords()
	if len(words) < 1000 {
		t.Fatal("Less than 1000 words loaded from dictionary")
	}
	if len(words) > 1000000 {
		t.Fatal("Less than 1000 words loaded from dictionary")
	}
}

func TestFindWords(t *testing.T) {
	t.Fatal("not implemented")
}
