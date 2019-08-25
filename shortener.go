package main

import (
	"errors"
	"math/rand"
	"strings"
)

// Database is a hash table of random letter keys and long URL values
type Database map[string]string

// ShortURLPrefix represents the domain of the URL shortener
var ShortURLPrefix = "www.short-url.com/"

// ErrNotFound is the error message for a URL not in the database
var ErrNotFound = errors.New("Short URL does not exist")

// ErrNotValidURL is the error message for a URL which does not meet the criteria
var ErrNotValidURL = errors.New("Not a valid URL")

// Decode retrieves the original URL by cross referenceing the random letter key
func (d Database) Decode(url string) (string, error) {
	key := strings.TrimPrefix(url, ShortURLPrefix)
	url, ok := d[key]

	if !ok {
		return "", ErrNotFound
	}
	return d[key], nil
}

// Encode creates a short URL and links it to the original long URL
func (d Database) Encode(url string) (string, error) {
	err := ValidateURL(url)
	key := RandStringBytes(6)
	d[key] = url
	shortURL := ShortURLPrefix + key
	return shortURL, err
}

// ValidateURL checks if the URL is valid
func ValidateURL(url string) error {
	if !strings.Contains(url, ".com") {
		return ErrNotValidURL
	}
	return nil
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// RandStringBytes creates a random string of 6 letters
// Source: https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func main() {}
