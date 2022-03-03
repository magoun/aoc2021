package sevenSegmentSearch

import (
	"aoc2021/helpers"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {

	dat, err := os.ReadFile("sevenSegmentSearch/input")
	// dat, err := os.ReadFile("sevenSegmentSearch/test")
	helpers.Check(err)

	stringDat := string(dat)
	stringArr := strings.Split(stringDat, "\n")

	patterns, outputs := parseSevenSegmentInput(stringArr)

	easyOccurences := get1478appearences(outputs)

	output := fmt.Sprintf("1, 4, 7, and 8 appear a total of %d times in the output digits.",
		easyOccurences)
	fmt.Println(output)

	digitsInBinarySegments := make(map[string]string)

	digitsInBinarySegments["1110111"] = "0"
	digitsInBinarySegments["0010010"] = "1"
	digitsInBinarySegments["1011101"] = "2"
	digitsInBinarySegments["1011011"] = "3"
	digitsInBinarySegments["0111010"] = "4"
	digitsInBinarySegments["1101011"] = "5"
	digitsInBinarySegments["1101111"] = "6"
	digitsInBinarySegments["1010010"] = "7"
	digitsInBinarySegments["1111111"] = "8"
	digitsInBinarySegments["1111011"] = "9"

	var results []string
	for index, output := range outputs {
		segmentDecoder := getSegmentWiringOrder(patterns[index])
		test := decodeOutputs(output, segmentDecoder, digitsInBinarySegments)
		results = append(results, test)
	}

	var total int
	for _, result := range results {
		intResult, _ := strconv.Atoi(result)
		total += intResult
	}

	fmt.Println(total)

}

func decodeOutputs(outputs []string, decoder [7]string, orderToDigitKey map[string]string) string {
	var results string

	for _, output := range outputs {
		digitString := ""

		for _, encodedLetter := range decoder {
			if strings.Contains(output, encodedLetter) {
				digitString += "1"
			} else {
				digitString += "0"
			}
		}

		decoded := orderToDigitKey[digitString]

		results += decoded
	}

	return results
}

func parseSevenSegmentInput(input []string) (patterns, outputs [][]string) {
	for _, line := range input {
		pieces := strings.Split(line, "|")
		patterns = append(patterns, strings.Split(pieces[0], " "))
		outputs = append(outputs, strings.Split(pieces[1], " "))
	}

	return
}

func get1478appearences(input [][]string) int {
	var total int

	for _, line := range input {
		for _, digit := range line {
			length := len(digit)
			switch length {
			case 2, 3, 4, 7:
				total++
			}
		}
	}

	return total
}

func getSegmentWiringOrder(patterns []string) [7]string {
	frequencies := make(map[string]int)
	var result [7]string
	var one, four string

	for _, pattern := range patterns {
		letters := strings.Split(pattern, "")

		if len(pattern) == 2 {
			one = pattern
		} else if len(pattern) == 3 {
			// seven = pattern
		} else if len(pattern) == 4 {
			four = pattern
		}

		for _, letter := range letters {
			frequencies[letter]++
		}
	}

	for letter, frequency := range frequencies {
		switch frequency {
		case 4:
			result[4] = letter
		case 6:
			result[1] = letter
		case 9:
			result[5] = letter
		case 8:
			if strings.Contains(one, letter) {
				result[2] = letter
			} else {
				result[0] = letter
			}
		case 7:
			if strings.Contains(four, letter) {
				result[3] = letter
			} else {
				result[6] = letter
			}
		}
	}

	return result
}
