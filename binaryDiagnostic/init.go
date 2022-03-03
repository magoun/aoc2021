package binaryDiagnostic

import (
	"aoc2021/helpers"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Run() {

	dat, err := os.ReadFile("binaryDiagnostic/input")
	// dat, err := os.ReadFile("binaryDiagnostic/test")
	helpers.Check(err)

	stringDat := string(dat)
	stringArr := strings.Split(stringDat, "\n")

	binLength := len(stringArr[0])

	binZeroCount := make([]int, binLength)
	total := 0

	for _, binary := range stringArr {

		binaryDigits := strings.Split(binary, "")

		for place, binDigit := range binaryDigits {
			if binDigit == "0" {
				binZeroCount[place]++
			}
		}

		total++

	}

	gammaRate := calcGammaRate(binZeroCount, binLength, total)
	epsilonRate := calcEpsilonRate(binZeroCount, binLength, total)

	output := fmt.Sprintf("Gamma: %d, Epsion: %d, Power: %d", gammaRate, epsilonRate, gammaRate*epsilonRate)
	fmt.Println(output)

	// Part 2
	oxygenGeneratorRating := getOxygenGeneratorRating(stringArr)
	co2ScrubberRating := getCO2ScrubberRating(stringArr)
	lifeSupportRating := oxygenGeneratorRating * co2ScrubberRating

	output = fmt.Sprintf("O2 Generator Rating: %d, CO2 Scrubber Rating: %d, Life Support Rating: %d", oxygenGeneratorRating, co2ScrubberRating, lifeSupportRating)
	fmt.Println(output)

}

func getCO2ScrubberRating(log []string) int {
	prefix := ""

	for len(log) > 1 {

		var oneLog, zeroLog []string
		checkDigitIndex := len(prefix)

		for _, binary := range log {

			checkDigit := string(binary[checkDigitIndex])

			if checkDigit == "0" {
				zeroLog = append(zeroLog, binary)
			} else {
				oneLog = append(oneLog, binary)
			}

		}

		if len(oneLog) >= len(zeroLog) {
			log = zeroLog
			prefix += "0"
		} else {
			log = oneLog
			prefix += "1"
		}

	}

	result, err := strconv.ParseInt(log[0], 2, 64)
	helpers.Check(err)

	return int(result)
}

func getOxygenGeneratorRating(log []string) int {
	prefix := ""

	for len(log) > 1 {

		var oneLog, zeroLog []string
		checkDigitIndex := len(prefix)

		for _, binary := range log {

			checkDigit := string(binary[checkDigitIndex])

			if checkDigit == "0" {
				zeroLog = append(zeroLog, binary)
			} else {
				oneLog = append(oneLog, binary)
			}

		}

		if len(zeroLog) > len(oneLog) {
			log = zeroLog
			prefix += "0"
		} else {
			log = oneLog
			prefix += "1"
		}

	}

	result, err := strconv.ParseInt(log[0], 2, 64)
	helpers.Check(err)

	return int(result)
}

func calcGammaRate(binZeroCount []int, binLength int, total int) int {
	mid := total / 2
	rate := 0
	shift := binLength - 1

	for place, zeroCount := range binZeroCount {
		if zeroCount < mid {
			rate += int(math.Pow(2, float64(shift-place)))
		}
	}

	return rate
}

func calcEpsilonRate(binZeroCount []int, binLength int, total int) int {
	mid := total / 2
	rate := 0
	shift := binLength - 1

	for place, zeroCount := range binZeroCount {
		if zeroCount > mid {
			rate += int(math.Pow(2, float64(shift-place)))
		}
	}

	return rate
}
