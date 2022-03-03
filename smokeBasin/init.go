package smokeBasin

import (
	"aoc2021/helpers"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Run() {

	dat, err := os.ReadFile("smokeBasin/input")
	// dat, err := os.ReadFile("smokeBasin/test")
	helpers.Check(err)

	stringDat := string(dat)
	stringArr := strings.Split(stringDat, "\n")

	sbMapArr := sbProcessInput(stringArr)
	sbNodes := makeNodes(sbMapArr)

	var totalRisk int

	for _, node := range sbNodes {
		totalRisk += node.getRiskLevel()
	}

	fmt.Println(totalRisk)

	basins := getBasinSizes(sbMapArr)
	sort.Sort(sort.Reverse(sort.IntSlice(basins)))

	basinCalc := basins[0] * basins[1] * basins[2]

	fmt.Println(basinCalc)
}

func getBasinSizes(sbMap [][]int) []int {
	height, width := getBasinDimensions(sbMap)
	var basins []int

	for h := 0; h <= height; h++ {
		for w := 0; w <= width; w++ {
			size := calcBasinSize(&sbMap, h, w, height, width)

			if size > 0 {
				basins = append(basins, size)
			}
		}
	}

	return basins
}

func calcBasinSize(sbMapRef *[][]int, hPos, wPos, hMax, wMax int) int {
	if (*sbMapRef)[hPos][wPos] == 9 {
		return 0
	}

	size := 1
	(*sbMapRef)[hPos][wPos] = 9

	if hPos > 0 {
		size += calcBasinSize(sbMapRef, hPos-1, wPos, hMax, wMax)
	}

	if hPos < hMax {
		size += calcBasinSize(sbMapRef, hPos+1, wPos, hMax, wMax)
	}

	if wPos > 0 {
		size += calcBasinSize(sbMapRef, hPos, wPos-1, hMax, wMax)
	}

	if wPos < wMax {
		size += calcBasinSize(sbMapRef, hPos, wPos+1, hMax, wMax)
	}

	return size
}

func getBasinDimensions(sbMap [][]int) (height, width int) {
	height = len(sbMap) - 1
	width = len(sbMap[0]) - 1
	return
}

type sbNode struct {
	north, south, east, west int
	height                   int
}

func (sbn *sbNode) getRiskLevel() int {
	if sbn.north <= sbn.height {
		return 0
	} else if sbn.south <= sbn.height {
		return 0
	} else if sbn.east <= sbn.height {
		return 0
	} else if sbn.west <= sbn.height {
		return 0
	}

	return sbn.height + 1
}

func sbProcessInput(input []string) [][]int {
	var result [][]int
	for _, line := range input {
		eastWestHeightsStr := strings.Split(line, "")
		var eastWestHeights []int
		for _, strDigit := range eastWestHeightsStr {
			intDigit, _ := strconv.Atoi(strDigit)
			eastWestHeights = append(eastWestHeights, intDigit)
		}
		result = append(result, eastWestHeights)
	}

	return result
}

func makeNodes(sbMap [][]int) []*sbNode {
	var nodes []*sbNode

	height := len(sbMap)
	width := len(sbMap[0])

	for southPos, values := range sbMap {
		for eastPos, value := range values {
			var northHeight, southHeight, eastHeight, westHeight int

			if southPos == 0 {
				northHeight = 9
			} else {
				northHeight = sbMap[southPos-1][eastPos]
			}

			if southPos == height-1 {
				southHeight = 9
			} else {
				southHeight = sbMap[southPos+1][eastPos]
			}

			if eastPos == 0 {
				westHeight = 9
			} else {
				westHeight = sbMap[southPos][eastPos-1]
			}

			if eastPos == width-1 {
				eastHeight = 9
			} else {
				eastHeight = sbMap[southPos][eastPos+1]
			}

			node := sbNode{
				north:  northHeight,
				south:  southHeight,
				east:   eastHeight,
				west:   westHeight,
				height: value}

			nodes = append(nodes, &node)
		}
	}

	return nodes
}
