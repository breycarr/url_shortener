package main

import (
	"errors"
	"strconv"
	"strings"
)

// ShortURLPrefix is represents the domain of the URL shortener
var ShortURLPrefix = "www.short-url.com/"

// Encode accepts a url string and returns a shortened url
func Encode(url string) (string, error) {
	if !strings.Contains(url, ".com") {
		return "", errors.New("Not a valid URL")
	}

	shortCode := 111111
	shortString := strconv.Itoa(shortCode)
	shortCode++
	return ShortURLPrefix + shortString, nil
}

func main() {}
