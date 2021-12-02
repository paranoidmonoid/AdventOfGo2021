package main

import (
	"log"
	"strconv"
)

func main() {
	lines, err := readLines("day01.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	result := 0
	for i := 0; i < len(lines)-1; i++ {
		current, _ := strconv.Atoi(lines[i])
		next, _ := strconv.Atoi(lines[i+1])
		if current < next {
			result++
		}
	}
	resultThree := 0
	for i := 0; i < len(lines)-3; i++ {
		first, _ := strconv.Atoi(lines[i])
		second, _ := strconv.Atoi(lines[i+1])
		third, _ := strconv.Atoi(lines[i+2])
		next, _ := strconv.Atoi(lines[i+3])
		currentThree := first + second + third
		if currentThree < (second + third + next) {
			resultThree++
		}
	}
	println(result)
	println(resultThree)
}
