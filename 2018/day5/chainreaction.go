package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

func main() {
	dat, _ := ioutil.ReadFile("input.txt")
	data := string(dat)
	min := len(data)
	for character := 1; character <= 26; character++ {
		lower := string(unicode.ToLower(rune('A' - 1 + character)))
		upper := string(rune('A' - 1 + character))
		polimers := strings.NewReplacer(lower, "", upper, "").Replace(data)
		polimers = react(polimers, 0)
		fmt.Println(len(polimers))
		if len(polimers) < min {
			min = len(polimers)
		}
	}
}

func react(polimers string, index int) string {
	if index == len(polimers)-1 {
		return polimers
	}

	charAtFirst := []rune(polimers)[index]
	charAtSecond := []rune(polimers)[index+1]

	if (unicode.IsLower(charAtFirst) && unicode.IsUpper(charAtSecond)) && unicode.ToUpper(charAtFirst) == charAtSecond {
		return replaceAtIndex(polimers, index)
	} else if (unicode.IsUpper(charAtFirst) && unicode.IsLower(charAtSecond)) && unicode.ToLower(charAtFirst) == charAtSecond {
		return replaceAtIndex(polimers, index)
	}

	return react(polimers, index+1)
}

func replaceAtIndex(str string, index int) string {
	str = str[:index] + str[index+2:]
	if index == 0 {
		return react(str, index)
	}
	return react(str, index-1)
}
