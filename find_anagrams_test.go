package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbcCalcualteWordProduct(t *testing.T) {
	primeMap := primeMap()

	abcProduct := calcualteWordProduct("abc", primeMap)
	cbaProduct := calcualteWordProduct("cba", primeMap)

	assert.Equal(t, abcProduct, 30, "word 'abc' product should equal 30 (2 * 3 * 5)")
	assert.Equal(t, abcProduct, cbaProduct, "'abc' and 'cba' products should be equal")
}

func TestPrimeMapFirstLetters(t *testing.T) {
	primeMap := primeMap()

	assert.Equal(t, primeMap[int('a')], 2, "prime map 'a' should equal 2")
	assert.Equal(t, primeMap[int('b')], 3, "prime map 'b' should equal 3")
	assert.Equal(t, primeMap[int('c')], 5, "prime map 'c' should equal 5")
}
func TestPrepare(t *testing.T) {
	primeMap := primeMap()
	anagrams := prepare("./test_data/abc.txt", primeMap)

	assert.Equal(t, len(anagrams[30]), 2)
	assert.Contains(t, anagrams[30], "abc")
	assert.Contains(t, anagrams[30], "cba")
}

func TestPrepareFull(t *testing.T) {
	primeMap := primeMap()
	anagrams := prepare("./test_data/lemmad.txt", primeMap)

	fmt.Println(anagrams)
}
