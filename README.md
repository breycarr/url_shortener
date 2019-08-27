# URL Shortener
> An API for a URL shortener 

## Specifications
- [x] Accepts a URL to be shortened.  
- [x] Generates a short URL for the original URL.  
- [x] Accepts a short URL and redirects the caller to the original URL.

#### Bonus points
- [x] Comes with a CLI that can be used to call your service.

## Installation
Clone this repository to your personal device.  
Navigate to the containing folder in the command line and run:  
```
go run shortener.go
```
You will see:
```
Please type a number then enter
1: Shorten URL
2: Get original URL
3: Exit program
```

Please follow the CLI instructions

## Usage
### Decode function
The Decode function accepts a Shortened URL as an argument:
* It validates that URL is currently present in the hash
    * If not it returns an error message 
* If the URL is valid, it returns the original URL

Input | Output
-|-
"www.tech.test/output" | "www.test-url.com/techtest/this_is_a_very_long_url/it_needs_to_be_much_shorter"
shortURL | Error message ("Short URL does not exist")

### Encode function
The Encode function accepts a URL string as an argument and performs several sub functions, which will be extracted in line with Single Responsibility
* It accepts a URL as an argument 
* It verifies that it is a valid URL
    * Defined as including either .com or .org as the most commonly used TLDs, but with scope for increasing this library
* It randomly generates a shortened URL
* It needs to verify that the shortened URL does not exist
    * If the URL does exist, 3 needs to occur again
* It saves that shortened URL as a key with the value of the original URL
* It returns that shortened URL as a string

Input | Output
-|-
"www.test-url.com/techtest/this_is_a_very_long_url/it_needs_to_be_much_shorter" | "www.tech.test/output"
testshortenthislongurlplease | Error Message ("Not a valid URL")
"" | Error Message ("Not a valid URL")

## Current limitations and possible solutions 
* The same long URL will receive different short URLs if given twice
    * With the use of key-value map, it is not possible to search through all the existant values to verify if it has already been added
* The database does not persist between sessions
    * the database map is created whenever shortener.go is run, so saved URLs do not persist between sessions

A possible solution to both of the above would be to implement database/SQL package, along with a method for checking the long URL against the database and retrieving and returning the matching short URL, this would function similarly to the loop used in GetUniqueKey.

* The returned URL does not open the window as a webpage

This could be implemented through use of the net/http package