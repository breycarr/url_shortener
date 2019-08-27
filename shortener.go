package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// Database is a hash table of letter keys and long URL values
type Database map[string]string

//  LengthOfKey represents the length of the short URL key
var lengthOfKey = 6

// ShortURLPrefix represents the domain of the URL shortener
var ShortURLPrefix = "www.short-url.com/"

// TopLevelDomains is a library of TLDs which are used to validate a URL
// The use of 2 TLDs is illustrative, more would be needed in a full implementation
var TopLevelDomains = []string{".com", ".org"}

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
	key := GetUniqueKey(d)
	d[key] = url
	shortURL := ShortURLPrefix + key
	return shortURL, err
}

// ValidateURL checks if the URL is valid
func ValidateURL(url string) error {
	for _, tld := range TopLevelDomains {
		if strings.Contains(url, tld) {
			return nil
		}
	}
	return ErrNotValidURL
}

// GetUniqueKey generates a random string and checks it is not present in the database
func GetUniqueKey(d Database) string {
	for {
		key := RandStringBytes(lengthOfKey)
		_, used := d[key]
		if used == false {
			return key
		}
	}
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

func main() {
	input := bufio.NewReader(os.Stdin)
	database := Database{}
	for {
		fmt.Println("Please type a number then enter")
		fmt.Println("1: Shorten URL")
		fmt.Println("2: Get original URL")
		fmt.Println("3: Exit program")
		cmdString, err := input.ReadString('\n')

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		switch cmdString {
		case "1":
			fmt.Println("Please enter the URL to shorten and type enter")
			longURL, _ := input.ReadString('\n')
			database.Encode(longURL)
		case "2":
			fmt.Println("Please enter your shortened URL and type enter")
			shortURL, _ := input.ReadString('\n')
			database.Decode(shortURL)
		case "3":
			break
		}
	}
}
