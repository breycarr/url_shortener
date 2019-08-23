package main

import "testing"

func TestEncode(t *testing.T) {
	t.Run("valid URL", func(t *testing.T) {
		url := "www.test-url.com/techtest/this_is_a_very_long_url/it_needs_to_be_much_shorter"

		got, _ := Encode(url)
		want := "www.short-url.com/111111"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("not a valid URL", func(t *testing.T) {
		badURL := "testshortenthislongurlplease"

		_, err := Encode(badURL)

		if err == nil {
			t.Error("wanted an error, but didn't get one")
		}
	})
}
