package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines, _ := readLines("day02/day02.txt")
	depth := 0
	horizontal := 0
	aim := 0
	// down forward up
	for _, line := range lines {
		command := strings.Split(line, " ")
		direction := command[0]
		value, _ := strconv.Atoi(command[1])
		if direction == "down" {
			aim += value
		} else if direction == "up" {
			aim -= value
		} else {
			horizontal += value
			depth += aim * value
		}
	}
	println(horizontal * depth)
}

// stackoverflow copypaste
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
