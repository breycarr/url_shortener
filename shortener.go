package main

import (
	"errors"
	"math/rand"
	"strings"
)

// ShortURLPrefix represents the domain of the URL shortener
var ShortURLPrefix = "www.short-url.com/"

// Database is a hash table of random letter keys and long URL values
type Database map[string]string

// Decode retrieves the original URL by cross referenceing the random letter key
func (d Database) Decode(url string) (string, error) {
	key := strings.TrimPrefix(url, ShortURLPrefix)
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
		return errors.New("Not a valid URL")
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
