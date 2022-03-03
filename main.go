package main

import (
	"aoc2021/sonarSweep"
)

func main() {
	sonarSweep.Run() // Day 1
	// Dive()                // Day 2
	// BinaryDiagnostic()    // Day 3
	// GiantSquid()          // Day 4
	// HydrothermalVenture() // Day 5
	// LanternFish()         // Day 6
	// TreacheryOfWhales()   // Day 7
	// SevenSegmentSearch() // Day 8
	// SmokeBasin()    // Day 9
	// SyntaxScoring() // Day 10
	// DumboOctopus()  // Day 11

}

func runeSlicePop(slice []rune) (rune, []rune) {
	length := len(slice)

	if length > 0 {
		popped := slice[length-1]
		slice = slice[:length-1]
		return popped, slice
	}

	panic(slice)
}
