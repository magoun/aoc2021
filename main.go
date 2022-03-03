package main

import (
	"aoc2021/binaryDiagnostic"
	"aoc2021/dive"
	"aoc2021/giantSquid"
	"aoc2021/hydrothermalVenture"
	"aoc2021/lanternFish"
	"aoc2021/sevenSegmentSearch"
	"aoc2021/smokeBasin"
	"aoc2021/sonarSweep"
	"aoc2021/syntaxScoring"
	"aoc2021/treacheryOfWhales"
)

func main() {
	sonarSweep.Run()          // Day 1
	dive.Run()                // Day 2
	binaryDiagnostic.Run()    // Day 3
	giantSquid.Run()          // Day 4
	hydrothermalVenture.Run() // Day 5
	lanternFish.Run()         // Day 6
	treacheryOfWhales.Run()   // Day 7
	sevenSegmentSearch.Run()  // Day 8
	smokeBasin.Run()          // Day 9
	syntaxScoring.Run()       // Day 10
	// dumboOctopus.Run()  // Day 11

}
