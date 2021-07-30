package main

import (
	"testing"
)

var Cases = []struct {
	in      string
	out     string
	isError bool
}{
	{
		in:      "nattu",
		out:     "neeweey",
		isError: false}}

func TestMain(t *testing.T) {
	for _, currentCase := range Cases {
		t.Run(currentCase.in, func(t *testing.T) {
			t.Fatalf("Test case failed for %v", currentCase)
		})
	}
}
