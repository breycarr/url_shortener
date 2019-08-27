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

//  letterBytes is used with the RandStringBytes function
// Source: https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

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
	if err != nil {
		return "", err
	}
	key := GetUniqueKey(d)
	d[key] = url
	shortURL := ShortURLPrefix + key
	return shortURL, err
}

// ValidateURL checks if the URL is valid
func ValidateURL(url string) error {
	validURL := false
	for _, tld := range TopLevelDomains {
		if strings.Contains(url, tld) {
			validURL = true
			break
		}
	}
	if !validURL {
		return ErrNotValidURL
	}
	return nil
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
		cmdString = strings.TrimSuffix(cmdString, "\n")

		switch cmdString {
		case "1":
			fmt.Println("\n Please enter the URL to shorten and type enter")
			longURL, _ := input.ReadString('\n')
			longURL = strings.TrimSuffix(longURL, "\n")
			output, err := database.Encode(longURL)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("\n Your shortened URL is: " + output + "\n")
			}
		case "2":
			fmt.Println("\n Please enter your shortened URL and type enter")
			shortURL, _ := input.ReadString('\n')
			shortURL = strings.TrimSuffix(shortURL, "\n")
			output, err := database.Decode(shortURL)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("\n Your original URL was: " + output + "\n")
			}
		case "3":
			fmt.Println("Goodbye")
			return
		}
	}
}
