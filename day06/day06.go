package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	println(calculateFishes(256, parseInput()))
}

func calculateFishes(days int, initialState []int) int {
	fishes := [9]int{}
	for _, fish := range initialState {
		fishes[fish]++
	}
	for i := 0; i < days; i++ {
		fishesBorn := fishes[0]
		for j := 1; j < 9; j++ {
			fishes[j-1] = fishes[j]
		}
		fishes[8] = fishesBorn
		fishes[6] += fishesBorn
	}
	sum := 0
	for _, n := range fishes {
		sum += n
	}
	return sum
}

func parseInput() []int {
	file, _ := os.Open("day06/day06.txt")
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	numbers := strings.Split(scanner.Text(), ",")
	var result []int
	for _, str := range numbers {
		n, _ := strconv.Atoi(str)
		result = append(result, n)
	}
	return result
}
