package dive

import (
	"aoc2021/helpers"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {

	dat, err := os.ReadFile("dive/input")
	// dat, err := os.ReadFile("dive/test")
	helpers.Check(err)

	stringDat := string(dat)
	stringArr := strings.Split(stringDat, "\n")

	horizontal := 0
	depth := 0

	for _, movement := range stringArr {

		movementSet := strings.Split(movement, " ")
		action := movementSet[0]
		unit, _ := strconv.Atoi(movementSet[1])

		switch action {
		case "forward":
			horizontal += unit
		case "down":
			depth += unit
		case "up":
			depth -= unit
		}

	}

	output := fmt.Sprintf("Horizontal: %d, Depth: %d, Multiplied: %d", horizontal, depth, horizontal*depth)
	fmt.Println(output)

	// Part 2
	horizontal = 0
	depth = 0
	aim := 0

	for _, movement := range stringArr {

		movementSet := strings.Split(movement, " ")
		action := movementSet[0]
		unit, _ := strconv.Atoi(movementSet[1])

		switch action {
		case "forward":
			horizontal += unit
			depth += aim * unit
		case "down":
			aim += unit
		case "up":
			aim -= unit
		}

	}

	output = fmt.Sprintf("Horizontal: %d, Depth: %d, Multiplied: %d", horizontal, depth, horizontal*depth)
	fmt.Println(output)

}
