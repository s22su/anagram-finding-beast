# Anagram finding beast

Anagram finding beast! Just give it a file & word and it will eat it
and spit out all the anagrams it finds for a given word.

**NB!** Dictionary file should be saved using UTF-8 encoding. `test_data/lemmad.txt` is converted to UTF-8.
[Original file](http://www.eki.ee/tarkvara/wordlist/lemmad.zip) had in Windows-1252 encoding.

## Usage

`./bin/find_anagrams_mac --dict ./test_data/lemmad.txt --word "kala"`
`./bin/find_anagrams_linux_amd64 --dict ./test_data/lemmad.txt --word "kala"`

**PS!** `./bin/find_anagrams_mac --dict ./test_data/lemmad.txt --word "mõõk"` works on Mac.
Not 100% sure if it works on Linux as my VM I uesd for building had some issues with õäöü.

## Output

`time_in_microseconds anagram1, anagram2, number_of_anagrams_found`

example:

`66695, kaal, 1`

## Dev requirements

You just need golang. Ver 12.0 was used for development
but most probably will work with 1.6+.

## Development

1. clone repo
1. `dep ensure`
1. `go test`
1. `go run find_anagrams.go --dict ./test_data/lemmad.txt --word kala`

## [Challenge](https://www.helmes.com/careers/challenge/)

Create an anagram finding beast and win a prize - Sony WH-1000XM3 wireless
headphones - the best noise cancelling headphones in the world.

## Desired outcome

The result should be an executable application, that finds all anagrams for provided word
from a provided dictionary as fast as possible. The application is responsible for calculating
the execution time and returning it in microseconds (μs). We will be testing your code against
the dictionary found here: http://www.eki.ee/tarkvara/wordlist/lemmad.zip
