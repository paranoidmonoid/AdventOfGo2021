package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	crabs := parseInput()
	median := median(crabs)
	fuel := 0
	for i := range crabs {
		f := crabs[i] - median
		if f < 0 {
			fuel += f * -1
		} else {
			fuel += f
		}
	}
	println(fuel)

	avg := avg(crabs)
	fuel2 := 0
	for i := range crabs {
		diff := crabs[i] - avg
		if diff < 0 {
			diff = diff * -1
		}
		fuel2 += ((1 + diff) * diff) / 2
	}
	println(fuel2)
}

func median(crabs []int) int {
	sort.Ints(crabs)
	n := len(crabs)
	if n%2 == 0 {
		return (crabs[n/2-1] + crabs[n/2]) / 2
	}
	return crabs[(n-1)/2]
}

func avg(crabs []int) int {
	count := 0
	for i := 0; i < len(crabs); i++ {
		count += crabs[i]
	}
	//round up
	if count%2 != 0 {
		count++
	}
	return count / len(crabs)
}

func parseInput() []int {
	file, _ := os.Open("day07/day07.txt")
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
