package main

import "strconv"

//func main() {
//	lines, _ := readLines("day03.txt")
//	firstPart03(lines)
//	secondPart03(lines)
//}

func secondPart03(lines []string) {
	resultMostCommon := findMostCommon(lines)
	resultLeastCommon := findLeastCommon(lines)
	println(resultMostCommon * resultLeastCommon)
}

func findMostCommon(lines []string) int64 {
	for i := 0; i < len(lines[0]); i++ {
		countZero := 0
		countOne := 0
		for j := 0; j < len(lines); j++ {
			if lines[j][i] == '0' {
				countZero++
			} else {
				countOne++
			}
		}
		temp := []string{}
		if countOne >= countZero {
			for k := 0; k < len(lines); k++ {
				if lines[k][i] == '1' {
					temp = append(temp, lines[k])
				}
			}
		} else {
			for k := 0; k < len(lines); k++ {
				if lines[k][i] == '0' {
					temp = append(temp, lines[k])
				}
			}
		}
		lines = temp
	}
	result, _ := strconv.ParseInt(lines[0], 2, 64)
	return result
}

func findLeastCommon(lines []string) int64 {
	for i := 0; i < len(lines[0]); i++ {
		if len(lines) == 1 {
			break
		}
		countZero := 0
		countOne := 0
		for j := 0; j < len(lines); j++ {
			if lines[j][i] == '0' {
				countZero++
			} else {
				countOne++
			}
		}
		temp := []string{}
		if countOne < countZero {
			for k := 0; k < len(lines); k++ {
				if lines[k][i] == '1' {
					temp = append(temp, lines[k])
				}
			}
		} else {
			for k := 0; k < len(lines); k++ {
				if lines[k][i] == '0' {
					temp = append(temp, lines[k])
				}
			}
		}
		lines = temp
	}
	result, _ := strconv.ParseInt(lines[0], 2, 64)
	return result
}

func firstPart03(lines []string) {
	resultGamma := ""
	resultEpsilon := ""
	for i := 0; i < len(lines[0]); i++ {
		countZero := 0
		countOne := 0
		for j := 0; j < len(lines); j++ {
			if lines[j][i] == '0' {
				countZero++
			} else {
				countOne++
			}
		}
		if countOne > countZero {
			resultGamma += "1"
			resultEpsilon += "0"
		} else {
			resultGamma += "0"
			resultEpsilon += "1"
		}
	}
	gamma, _ := strconv.ParseInt(resultGamma, 2, 64)
	epsilon, _ := strconv.ParseInt(resultEpsilon, 2, 64)
	println(gamma * epsilon)
}
