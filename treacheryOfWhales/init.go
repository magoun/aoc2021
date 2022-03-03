package treacheryOfWhales

import (
	"aoc2021/helpers"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {

	dat, err := os.ReadFile("treacheryOfWhales/input")
	// dat, err := os.ReadFile("treacheryOfWhales/test")
	helpers.Check(err)

	stringDat := string(dat)
	subPositionsRaw := strings.Split(stringDat, ",")

	var subPositions []int

	for _, stringVal := range subPositionsRaw {
		intVal, _ := strconv.Atoi(stringVal)
		subPositions = append(subPositions, intVal)
	}

	idealPosition, requiredFuel := getOptimalMovement(subPositions)

	output := fmt.Sprintf("The ideal horizonal alignment is position %d, requiring %d fuel.", idealPosition, requiredFuel)
	fmt.Println(output)

}

func calculateFuel(positions []int, destination int) int {
	var fuel int

	for _, start := range positions {
		if start > destination {
			// fuel += start - destination
			fuel += sumToNumber(start - destination)
		} else {
			// fuel += destination - start
			fuel += sumToNumber(destination - start)
		}
	}

	return fuel
}

func sumToNumber(n int) int {
	return n * (n + 1) / 2
}

func getOptimalMovement(positions []int) (idealPosition, requiredFuel int) {
	min := positions[0]
	var max int

	for _, value := range positions {
		if value < min {
			min = value
		}

		if value > max {
			max = value
		}
	}

	requiredFuel = calculateFuel(positions, 0)

	for i := min; i <= max; i++ {
		fuel := calculateFuel(positions, i)

		if fuel < requiredFuel {
			requiredFuel = fuel
			idealPosition = i
		}
	}

	return
}
