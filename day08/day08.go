package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	_, outputs := parseInput()
	println(countSimple(outputs))
	//sum := 0
	//for i := 0; i < len(references); i++ {
	//	sum += parseSignals(references[i], outputs[i])
	//}
	//println(sum)
}

func parseInput() ([][]string, [][]string) {
	file, _ := os.Open("day08/day08.txt")
	var references [][]string
	var outputs [][]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " | ")
		r := strings.Split(line[0], " ")
		o := strings.Split(line[1], " ")
		references = append(references, r)
		outputs = append(outputs, o)
	}

	return references, outputs
}

func countSimple(outputs [][]string) int {
	count := 0
	for i := range outputs {
		for j := 0; j < len(outputs[0]); j++ {
			if len(outputs[i][j]) == 2 || len(outputs[i][j]) == 4 || len(outputs[i][j]) == 3 || len(outputs[i][j]) == 7 {
				count++
			}
		}
	}
	return count
}

//func parseSignals(reference []string, output []string) int {
//
//	var frequencies [10]int64
//	for i := range reference {
//		for _, ch := range reference[i] {
//			frequencies[i] += int64(ch) - 'a'
//		}
//	}
//	res := numberToDigit(output[0], frequencies)*1000 + numberToDigit(output[1], frequencies)*100 +
//		numberToDigit(output[2], frequencies)*10 + numberToDigit(output[3], frequencies)
//	println(res)
//	return res
//
//}
//
//func numberToDigit(number string, frequencies [10]int64) int {
//	length := len(number)
//	if length == 2 {
//		return 1
//	}
//	if length == 4 {
//		return 4
//	}
//	if length == 3 {
//		return 7
//	}
//	if length == 7 {
//		return 8
//	}
//	if length == 6 {
//		println("*******")
//		println(number)
//		println(frequencies[number[0]-'a'])
//		println(frequencies[number[1]-'a'])
//		println(frequencies[number[2]-'a'])
//		println(frequencies[number[3]-'a'])
//		println(frequencies[number[4]-'a'])
//		println(frequencies[number[5]-'a'])
//		println("*******")
//		if frequencies[number[2]-'a'] == 8 {
//			return 0
//		}
//		if frequencies[number[3]-'a'] == 4 {
//			return 6
//		}
//		return 9
//	}
//	// length == 5
//	if frequencies[number[1]-'a'] == 6 {
//		return 5
//	}
//	if frequencies[number[3]-'a'] == 4 {
//		return 2
//	}
//	return 3
//}
//
////wrongs: 6 (writes 9) 2 (writes 3) 5 (writes 3)
