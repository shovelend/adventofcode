package main

import (
	"fmt"
	"io/ioutil"
)

type Guard struct {
	id     string
	startX int
	startY int
	width  int
	length int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("input.txt")
	check(err)
	fmt.Println(dat)
}
