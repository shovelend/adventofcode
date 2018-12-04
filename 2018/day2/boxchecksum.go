package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("input.txt")
	check(err)
	boxIds := strings.Split(string(dat), "\n")
	var twos int
	var threes int
	var checksum int

	for _, boxID := range boxIds {
		two := doesContainXOfTheSameLetters(boxID, 2)
		three := doesContainXOfTheSameLetters(boxID, 3)

		if two {
			twos++
		}
		if three {
			threes++
		}
	}
	checksum = twos * threes
	fmt.Println(checksum)
}

func doesContainXOfTheSameLetters(id string, nOfLetters int) bool {
	symbols := frequentSym(id)
	for _, v := range symbols {
		if v == nOfLetters {
			return true
		}
	}
	return false
}

func frequentSym(id string) map[string]int {
	symbols := make(map[string]int)

	for _, sym := range id {
		char := strconv.QuoteRune(sym)
		count := symbols[char]

		if count == symbols[char] {
			symbols[char] = count + 1
		} else {
			symbols[char] = 1
		}
	}

	return symbols
}
