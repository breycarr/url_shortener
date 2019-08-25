# URL Shortener
> An API for a URL shortener 

## Specifications
- [x] Accepts a URL to be shortened.  
- [x] Generates a short URL for the original URL.  
- [] Accepts a short URL and redirects the caller to the original URL.

#### Bonus points
- [] Comes with a CLI that can be used to call your service.

## Installation

## Usage
### Decode function
The Decode function accepts a Shortened URL as an argument:
1. It validates that URL is currently present in the hash <ul><li>- [x]</li></ul>
    * If not it returns an error message <ul><li>- [x]</li></ul>
2. If the URL is valid, it returns the original URL <ul><li>- [x]</li></ul>
3. It redirects the caller to the original URL <ul><li>- []</li></ul>

Input | Output
-|-
"www.tech.test/output" | "www.test-url.com/techtest/this_is_a_very_long_url/it_needs_to_be_much_shorter"
shortURL | Error message ("Short URL does not exist")

### Encode function
The Encode function accepts a URL string as an argument and performs several sub functions, which will be extracted in line with Single Responsibility
1. It accepts a URL as an argument <ul><li>- [x]</li></ul>
2. It verifies that it is a valid URL <ul><li>- []</li></ul>
    * Defined as including either .com or .org as the most commonly used TLDs, but with scope for increasing this library
3. It randomly generates a shortened URL <ul><li>- [x]</li></ul>
4. It needs to verify that the shortened URL does not exist <ul><li>- [x]</li></ul>
    * If the URL does exist, 3 needs to occur again
5. It saves that shortened URL as a key with the value of the original URL <ul><li>- [x]</li></ul>
6. it returns that shortened URL as a string <ul><li>- [x]</li></ul>

Input | Output
-|-
"www.test-url.com/techtest/this_is_a_very_long_url/it_needs_to_be_much_shorter" | "www.tech.test/output"
testshortenthislongurlplease | Error Message ("Not a valid URL")
"" | Error Message ("Not a valid URL")

### Edge Cases
* The same URL is given twice

