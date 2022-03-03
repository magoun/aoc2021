package hydrothermalVenture

import (
	"aoc2021/helpers"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {

	dat, err := os.ReadFile("hydrothermalVenture/input")
	// dat, err := os.ReadFile("hydrothermalVenture/test")
	helpers.Check(err)

	stringDat := string(dat)
	stringArr := strings.Split(stringDat, "\n")

	ventLines := getVentLinesFromInput(stringArr)

	filledVentMap := getVentMapVorH(ventLines)
	hotSpots := countHotSpots(filledVentMap)

	filledVentMapWithDiagonals := addDiagonalsToVentMap(filledVentMap, ventLines)
	hotSpotsWithDiagonals := countHotSpots(filledVentMapWithDiagonals)

	output := fmt.Sprintf("There are %d vent overlap points without diagonals, and %d including diagonals.", hotSpots, hotSpotsWithDiagonals)
	fmt.Println(output)

}

func addDiagonalsToVentMap(ventMap [][]int, ventLines []*VentLine) [][]int {
	for _, vl := range ventLines {
		if vl.isDiagonal() {
			begin_x := vl.start.x
			end_x := vl.end.x
			begin_y := vl.start.y

			if vl.start.x > vl.end.x {
				begin_x = vl.end.x
				end_x = vl.start.x
				begin_y = vl.end.y
			}

			slope, _ := vl.slope()

			for x := begin_x; x <= end_x; x++ {
				ventMap[x][(slope*(x-begin_x)+begin_y)]++
			}
		}
	}

	return ventMap
}

func countHotSpots(ventMap [][]int) int {
	var count int

	for _, rows := range ventMap {
		for _, value := range rows {
			if value > 1 {
				count++
			}
		}
	}

	return count
}

func getVentMapVorH(ventLines []*VentLine) [][]int {
	max_x, max_y := getLargestDimensionsVorH(ventLines)
	emptyVentMap := getEmptyVentMap(max_x, max_y)

	for _, vl := range ventLines {
		if vl.isHorizontal() {
			begin_x := vl.start.x
			end_x := vl.end.x

			if vl.start.x > vl.end.x {
				begin_x = vl.end.x
				end_x = vl.start.x
			}

			for x := begin_x; x <= end_x; x++ {
				emptyVentMap[x][vl.start.y]++
			}
		} else if vl.isVertical() {
			begin_y := vl.start.y
			end_y := vl.end.y

			if vl.start.y > vl.end.y {
				begin_y = vl.end.y
				end_y = vl.start.y
			}

			for y := begin_y; y <= end_y; y++ {
				emptyVentMap[vl.start.x][y]++
			}
		}
	}

	return emptyVentMap
}

func getEmptyVentMap(max_x, max_y int) [][]int {
	var emptyMap [][]int
	var emptyY []int

	for y := 0; y <= max_y; y++ {
		emptyY = append(emptyY, 0)
	}

	for x := 0; x <= max_x; x++ {
		newX := make([]int, len(emptyY))
		copy(newX, emptyY)
		emptyMap = append(emptyMap, newX)
	}

	return emptyMap
}

func getLargestDimensionsVorH(ventLines []*VentLine) (max_x, max_y int) {
	for _, vl := range ventLines {
		if vl.isHorizontal() || vl.isVertical() {
			if vl.start.x > max_x {
				max_x = vl.start.x
			}
			if vl.end.x > max_x {
				max_x = vl.end.x
			}
			if vl.start.y > max_y {
				max_y = vl.start.y
			}
			if vl.end.y > max_y {
				max_y = vl.end.y
			}
		}
	}

	return
}

func getVentLinesFromInput(input []string) (ventLines []*VentLine) {
	for _, inputLine := range input {
		points := strings.Split(inputLine, " -> ")

		start := newPoint(points[0])
		end := newPoint(points[1])
		// fmt.Println(start, end)

		ventLines = append(ventLines, newVentLine(start, end))
	}

	return
}

type VentLine struct {
	start *Point
	end   *Point
}

// func (vl *VentLine) getPoints() []*Point {

// }

func newVentLine(start, end *Point) *VentLine {
	vl := VentLine{start: start, end: end}
	return &vl
}

type Point struct {
	x int
	y int
}

func newPoint(xy string) *Point {
	xyArr := strings.Split(xy, ",")

	x, _ := strconv.Atoi(xyArr[0])
	y, _ := strconv.Atoi(xyArr[1])

	point := Point{x: x, y: y}

	return &point
}

func (vl *VentLine) isVertical() bool {
	return vl.start.x == vl.end.x
}

func (vl *VentLine) isHorizontal() bool {
	return vl.start.y == vl.end.y
}

func (vl *VentLine) isDiagonal() bool {
	slope, ok := vl.slope()

	if !ok {
		return false
	}

	return slope == 1 || slope == -1
}

func (vl *VentLine) slope() (int, bool) {
	if vl.end.x == vl.start.x {
		return 0, false
	}

	return (vl.end.y - vl.start.y) / (vl.end.x - vl.start.x), true
}
