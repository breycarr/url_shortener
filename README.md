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
Navigate to the containing folder in the command line and run
`go run shortener.go`

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

### Edge Cases
* The same URL is given twice

