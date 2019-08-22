package main

import "testing"

func TestEncode(t *testing.T) {
	url := "www.test-url.com/techtest/"

	got := Encode(url)
	want := "output"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
