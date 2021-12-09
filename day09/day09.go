package main

import (
	"bufio"
	"os"
)

func main() {
	layout := readCavesLayout()
	println(calculateRisk(layout))
}

func readCavesLayout() [][]int {
	file, _ := os.Open("day09/day09.txt")
	var layout [][]int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		var numbers []int
		for _, ch := range line {
			numbers = append(numbers, int(ch-'0'))
		}
		layout = append(layout, numbers)
	}
	return layout
}

func calculateRisk(layout [][]int) int {
	risk := 0
	for i := 0; i < len(layout); i++ {
		for j := 0; j < len(layout[0]); j++ {
			// current element: (i, j)
			// neighbours: (i+1, j), (i-1, j), (i, j+1), (i, j-1)
			current := layout[i][j]
			var neighbours []int
			if i != 0 {
				neighbours = append(neighbours, layout[i-1][j])
			}
			if j != 0 {
				neighbours = append(neighbours, layout[i][j-1])
			}
			if i != len(layout)-1 {
				neighbours = append(neighbours, layout[i+1][j])
			}
			if j != len(layout[0])-1 {
				neighbours = append(neighbours, layout[i][j+1])
			}
			currentRisk := current + 1
			for _, neighbour := range neighbours {
				if current >= neighbour {
					currentRisk = 0
					break
				}
			}
			risk += currentRisk
		}
	}
	return risk
}
