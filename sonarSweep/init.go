package sonarSweep

import (
	"aoc2021/helpers"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {

	dat, err := os.ReadFile("sonarSweep/input")
	helpers.Check(err)

	stringDat := string(dat)
	stringArr := strings.Split(stringDat, "\n")

	prevDepth := 0
	increases := 0

	for _, depth := range stringArr {

		intDepth, _ := strconv.Atoi(depth)

		if intDepth > prevDepth {
			increases++
		}

		prevDepth = intDepth
	}

	fmt.Println(increases - 1) // First increase doesn't count, since it isn't "increasing" from the initialization of 0

	var avgDepths []int

	currentDepths := [3]int{0, 0, 0}

	for _, depth := range stringArr {

		intDepth, _ := strconv.Atoi(depth)

		currentDepths[0] = currentDepths[1]
		currentDepths[1] = currentDepths[2]
		currentDepths[2] = intDepth

		if currentDepths[0] > 0 && currentDepths[1] > 0 && currentDepths[2] > 0 {
			currentAvg := currentDepths[0] + currentDepths[1] + currentDepths[2]
			avgDepths = append(avgDepths, currentAvg)
		}

	}

	prevAvgDepth := 0
	avgDepthIncreases := -1

	for _, avgDepth := range avgDepths {

		if avgDepth > prevAvgDepth {
			avgDepthIncreases++
		}

		prevAvgDepth = avgDepth
	}

	fmt.Println(avgDepthIncreases)

}
