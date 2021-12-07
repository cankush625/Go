package main

import (
	"reflect"
	"testing"
)

func TestWordCount(t *testing.T) {
	cases := []struct{
		word string
		res map[string]int
	}{
		{"Hello World", map[string]int{"Hello": 1, "World": 1}},
		{"", map[string]int{}},
	}
	for _, c := range cases {
		res := WordCount(c.word)
		if !reflect.DeepEqual(res, c.res) {
			t.Errorf("Word count of %v should be %v", c.word, c.res)
		}
	}
}
