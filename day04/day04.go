package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	numbers, bingos := parseInput()
	println(play(numbers, bingos))
	println(playlong(numbers, bingos))
}

func play(numbers []int, bingos [][][]int) int {

	for _, n := range numbers {
		for _, bingo := range bingos {
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					if bingo[i][j] == n {
						bingo[i][j] = -1
						if checkBingo(bingo, i, j) {
							return sumBingo(bingo) * n
						}
					}
				}
			}
		}
	}
	return -1
}

func playlong(numbers []int, bingos [][][]int) int {
	sum := 0
	for _, n := range numbers {
		var temp [][][]int
		for _, bingo := range bingos {
			interesting := false
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					if bingo[i][j] == n {
						bingo[i][j] = -1
						if checkBingo(bingo, i, j) {
							sum = sumBingo(bingo) * n
							interesting = true
						}
					}
				}
			}
			if !interesting {
				temp = append(temp, bingo)
			}
		}
		bingos = temp
	}
	return sum
}

func sumBingo(bingo [][]int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if bingo[i][j] != -1 {
				sum += bingo[i][j]
			}
		}
	}
	return sum
}

func checkBingo(bingo [][]int, v int, h int) bool {
	line := true
	column := true
	for i := 0; i < 5; i++ {
		if bingo[v][i] != -1 {
			line = false
		}
	}
	for i := 0; i < 5; i++ {
		if bingo[i][h] != -1 {
			column = false
		}
	}
	return line || column
}

func parseInput() ([]int, [][][]int) {
	file, _ := os.Open("day04/day04.txt")
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	numbersStr := strings.Split(scanner.Text(), ",")
	var numbers []int
	for _, n := range numbersStr {
		current, _ := strconv.Atoi(n)
		numbers = append(numbers, current)
	}

	var bingos [][][]int
	re := regexp.MustCompile("\\s+")
	for scanner.Scan() {
		current := scanner.Text()
		if current != "\n" {
			var bingo [][]int
			for i := 0; i < 5; i++ {
				scanner.Scan()
				str := re.Split(strings.TrimSpace(scanner.Text()), 5)
				var line []int
				for _, n := range str {
					lineNumber, _ := strconv.Atoi(n)
					line = append(line, lineNumber)
				}
				bingo = append(bingo, line)
			}
			bingos = append(bingos, bingo)
		}
	}
	return numbers, bingos
}
