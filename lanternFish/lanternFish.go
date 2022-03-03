package lanternFish

import (
	"aoc2021/helpers"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {

	dat, err := os.ReadFile("lanternFish/input")
	// dat, err := os.ReadFile("lanternFish/test")
	helpers.Check(err)

	stringDat := string(dat)
	// stringArr := strings.Split(stringDat, "\n")
	fishClocksStr := strings.Split(stringDat, ",")
	var fishClocks []int

	for _, stringVal := range fishClocksStr {
		intVal, _ := strconv.Atoi(stringVal)
		fishClocks = append(fishClocks, intVal)
	}

	fishPop := newFishPop(fishClocks)

	days := 256

	for day := 1; day <= days; day++ {
		fishPop.modelDay()
	}

	output := fmt.Sprintf("There are a total of %d fish after %d days.", fishPop.getTotal(), days)
	fmt.Println(output)

}

type FishPop struct {
	spawningDays [9]int
}

func newFishPop(initialPop []int) *FishPop {
	var pop [9]int

	for _, days := range initialPop {
		pop[days]++
	}

	fp := FishPop{pop}

	return &fp
}

func (fp *FishPop) modelDay() {
	var newSpawningDays [9]int

	for day, fish := range fp.spawningDays {
		if day > 0 {
			newSpawningDays[day-1] = fish
		}
	}

	newSpawningDays[8] = fp.spawningDays[0]
	newSpawningDays[6] += fp.spawningDays[0]

	fp.spawningDays = newSpawningDays
}

func (fp *FishPop) getTotal() int {
	total := 0

	for _, fish := range fp.spawningDays {
		total += fish
	}

	return total
}
