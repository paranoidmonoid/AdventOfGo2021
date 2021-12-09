package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func main() {
	input := parseInput()
	println(filterSimpleLines(input))
	println(calculateField(input))
}

func parseInput() [][]int {
	file, _ := os.Open("day05/day05.txt")
	var coordinates [][]int
	re := regexp.MustCompile("(,| -> )")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		splitToPoints := re.Split(scanner.Text(), 4)
		var points []int
		for i := 0; i < 4; i++ {
			point, _ := strconv.Atoi(splitToPoints[i])
			points = append(points, point)
		}
		coordinates = append(coordinates, points)
	}
	return coordinates
}

func filterSimpleLines(coordinates [][]int) int {
	maxX := 0
	maxY := 0
	var simpleLines [][]int
	for _, coordinate := range coordinates {
		if coordinate[0] > maxX {
			maxX = coordinate[0]
		}
		if coordinate[2] > maxX {
			maxX = coordinate[2]
		}
		if coordinate[1] > maxY {
			maxY = coordinate[1]
		}
		if coordinate[3] > maxY {
			maxY = coordinate[3]
		}
		if coordinate[0] == coordinate[2] || coordinate[1] == coordinate[3] {
			simpleLines = append(simpleLines, coordinate)
		}
	}
	return calculateSimpleField(simpleLines, maxX, maxY)
}

func calculateSimpleField(coordinates [][]int, maxX int, maxY int) int {
	var field = make([][]int, maxX+1)
	for i := range field {
		field[i] = make([]int, maxY+1)
	}
	for _, coordinate := range coordinates {
		if coordinate[0] == coordinate[2] {
			x := coordinate[0]
			if coordinate[1] < coordinate[3] {
				for y := coordinate[1]; y <= coordinate[3]; y++ {
					field[x][y]++
				}
			} else {
				for y := coordinate[3]; y <= coordinate[1]; y++ {
					field[x][y]++
				}
			}
		} else {
			y := coordinate[1]
			if coordinate[0] < coordinate[2] {
				for x := coordinate[0]; x <= coordinate[2]; x++ {
					field[x][y]++
				}
			} else {
				for x := coordinate[2]; x <= coordinate[0]; x++ {
					field[x][y]++
				}
			}
		}
	}
	count := 0
	for _, row := range field {
		for _, point := range row {
			if point > 1 {
				count++
			}
		}
	}
	return count
}

func calculateField(coordinates [][]int) int {
	maxX := 0
	maxY := 0
	for _, coordinate := range coordinates {
		if coordinate[0] > maxX {
			maxX = coordinate[0]
		}
		if coordinate[2] > maxX {
			maxX = coordinate[2]
		}
		if coordinate[1] > maxY {
			maxY = coordinate[1]
		}
		if coordinate[3] > maxY {
			maxY = coordinate[3]
		}
	}
	var field = make([][]int, maxX+1)
	for i := range field {
		field[i] = make([]int, maxY+1)
	}
	for _, coordinate := range coordinates {
		x1 := coordinate[0]
		x2 := coordinate[2]
		y1 := coordinate[1]
		y2 := coordinate[3]
		if x1 == x2 {
			if y1 < y2 {
				for y := y1; y <= y2; y++ {
					field[x1][y]++
				}
			} else {
				for y := y2; y <= y1; y++ {
					field[x1][y]++
				}
			}
		} else if y1 == y2 {
			if x1 < x2 {
				for x := x1; x <= x2; x++ {
					field[x][y1]++
				}
			} else {
				for x := x2; x <= x1; x++ {
					field[x][y1]++
				}
			}
		} else if x1 > x2 {
			if y1 > y2 {
				for i := 0; i <= (x1 - x2); i++ {
					field[x2+i][y2+i]++
				}
			} else {
				for i := 0; i <= (x1 - x2); i++ {
					field[x2+i][y2-i]++
				}
			}
		} else if x1 < x2 {
			if y1 > y2 {
				for i := 0; i <= (x2 - x1); i++ {
					field[x1+i][y1-i]++
				}
			} else {
				for i := 0; i <= (x2 - x1); i++ {
					field[x1+i][y1+i]++
				}
			}
		}
	}
	count := 0
	for _, row := range field {
		for _, point := range row {
			if point > 1 {
				count++
			}
		}
	}
	return count
}
