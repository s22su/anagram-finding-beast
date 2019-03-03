package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	startTime := time.Now()

	word := flag.String("word", "", "word to search anagrams for")
	filename := flag.String("dict", "", "path for dictionary file location")

	flag.Parse()

	if len(*filename) == 0 {
		fmt.Println("Please provide distionary file path with --dict flag")
		os.Exit(1)
	}

	if len(*word) == 0 {
		fmt.Println("Please provide word with --word flag")
		os.Exit(1)
	}

	primeMap := primeMap()
	anagrams := prepare(*filename, primeMap)

	wp := calcualteWordProduct(*word, primeMap)
	foundAnagrams := anagrams[wp]

	// fastest version to remove initial word
	for i, v := range foundAnagrams {
		if v == *word {
			foundAnagrams[i] = foundAnagrams[len(foundAnagrams)-1]
			foundAnagrams[len(foundAnagrams)-1] = ""
			foundAnagrams = foundAnagrams[:len(foundAnagrams)-1]
			break
		}
	}

	endTime := time.Now()

	totalTimeInMicroSeconds := endTime.Sub(startTime) / time.Microsecond

	fmt.Printf("%d, %s, %d\n", totalTimeInMicroSeconds, strings.Join(foundAnagrams, ", "), len(foundAnagrams))
}

func prepare(filename string, primeMap map[int]int) map[int][]string {
	anagrams := make(map[int][]string)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		product := calcualteWordProduct(word, primeMap)

		anagrams[product] = append(anagrams[product], word)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return anagrams
}

// Map each char of word to a prime number and multiply all prime numbers
func calcualteWordProduct(word string, primeMap map[int]int) int {
	p := 1
	for _, r := range word {
		// NB! Int overflow, but GO is smart enough and as
		// it is not really needed to know exact number, its ok :)
		p = p * primeMap[int(r)]
	}

	return p
}

// Map each char to a prime number
func primeMap() map[int]int {
	// here are some spare primes if needed :)
	// 163, 167, 173, 179, 181, 191, 193, 197, 199

	primeMap := map[int]int{
		int('a'):  2,
		int('b'):  3,
		int('c'):  5,
		int('-'):  7,
		int('\''): 11,
		int('s'):  13,
		int('q'):  17,
		int('f'):  19,
		int('x'):  23,
		int('k'):  29,
		int('j'):  31,
		int('v'):  37,
		int('z'):  41,
		int('r'):  43,
		int('ö'):  47,
		int('p'):  53,
		int('t'):  59,
		int('ä'):  61,
		int('é'):  67,
		int('m'):  71,
		int('h'):  73,
		int('y'):  79,
		int('w'):  83,
		int('i'):  89,
		int('ð'):  97,
		int('þ'):  101,
		int('l'):  103,
		int('ü'):  107,
		int('õ'):  109,
		int('e'):  113,
		int('g'):  127,
		int('n'):  131,
		int(' '):  137,
		int('o'):  139,
		int('!'):  149,
		int('d'):  151,
		int('u'):  157,
	}

	return primeMap
}

// Helper to get all chars from file that
// need to be mapped to prime
func collectCharsFromFile(filename string) {
	chars := make(map[string]string)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())

		for _, r := range word {
			currentChar := string(r)
			if _, ok := chars[currentChar]; !ok {
				chars[currentChar] = currentChar
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, char := range chars {
		fmt.Printf("int('%s'): %s,\n", char, "X")
	}
}
