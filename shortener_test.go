package main

import (
	"regexp"
	"testing"
)

var longURL = "www.test-url.com/techtest/this_is_a_very_long_url/it_needs_to_be_much_shorter"
var shortURL = "www.short-url.com/abcdef"
var key = "abcdef"
var shortURLregexp = `www.short-url.com\/[A-Za-z]{6}`

func TestDecode(t *testing.T) {
	t.Run("existing short URL", func(t *testing.T) {
		database := Database{key: longURL}
		got, err := database.Decode(shortURL)
		want := longURL

		AssertNoError(t, err)
		AssertStrings(t, got, want)
	})

	t.Run("non-existant short URL", func(t *testing.T) {
		database := Database{}
		_, err := database.Decode(shortURL)

		AssertError(t, err)
	})
}

func TestEncode(t *testing.T) {
	database := Database{}

	t.Run("shortens URL", func(t *testing.T) {
		got, err := database.Encode(longURL)
		want, _ := regexp.MatchString(shortURLregexp, got)

		AssertNoError(t, err)
		if !want {
			t.Errorf("got %q", got)
		}
	})

	t.Run("stores URL", func(t *testing.T) {
		output, err := database.Encode(longURL)
		got, _ := database.Decode(output)
		want := longURL

		AssertNoError(t, err)
		AssertStrings(t, got, want)
	})

	t.Run("not a valid URL", func(t *testing.T) {
		badURL := "testshortenthislongurlplease"
		_, err := database.Encode(badURL)

		AssertError(t, err)
	})
}

func AssertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func AssertError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Fatal("wanted an error, but didn't get one")
	}
}

func AssertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("Unexpected error encountered")
	}
}
