package syntaxScoring

import (
	"aoc2021/helpers"
	"fmt"
	"os"
	"sort"
	"strings"
)

func Run() {

	dat, err := os.ReadFile("syntaxScoring/input")
	// dat, err := os.ReadFile("syntaxScoring/test")
	helpers.Check(err)

	stringDat := string(dat)
	stringArr := strings.Split(stringDat, "\n")

	incomplete, corrupted := parseSyntaxInput(stringArr)

	points := sumCorrupted(corrupted)

	output := fmt.Sprintf("Day 10 - Total for corrupted syntax errors: %d", points)
	fmt.Println(output)

	points = getIncompleteScore(incomplete)

	output = fmt.Sprintf("Day 10 - Middle score for incomplete syntax errors: %d", points)
	fmt.Println(output)

}

func getIncompleteScore(incomplete []string) int {

	closeStacks := getCloseStacks(incomplete)

	var stackScores []int

	for _, stack := range closeStacks {
		var score int

		// closeStacks are revered, so iterate backwards
		for i := len(stack) - 1; i >= 0; i-- {
			score *= 5
			score += getCloserPointValue(stack[i])
		}

		stackScores = append(stackScores, score)
	}

	mid := (len(stackScores) - 1) / 2

	sort.Ints(stackScores)

	return stackScores[mid]
}

func getCloserPointValue(closer rune) int {
	switch closer {
	case ')':
		return 1
	case '}':
		return 3
	case ']':
		return 2
	case '>':
		return 4
	default:
		panic(closer)
	}
}

func getCloseStacks(incomplete []string) (closeStacks [][]rune) {
	for _, openStack := range incomplete {
		var closeStack []rune

		for _, char := range openStack {
			closeStack = append(closeStack, pairCloser(char))
		}

		closeStacks = append(closeStacks, closeStack)
	}

	return
}

func sumCorrupted(corrupted []rune) int {
	var total int

	for _, char := range corrupted {
		switch char {
		case ')':
			total += 3
		case ']':
			total += 57
		case '}':
			total += 1197
		case '>':
			total += 25137
		}
	}

	return total
}

func parseSyntaxInput(input []string) (incomplete []string, corrupted []rune) {
	openers := "([{<"
	var popped rune

	for _, line := range input {
		isCorrupted := false
		openStack := make([]rune, 0)

		for _, char := range line {

			if strings.ContainsRune(openers, char) {
				openStack = append(openStack, char)
			} else {
				popped, openStack = helpers.RuneSlicePop(openStack)

				if pairOpener(char) != popped {
					corrupted = append(corrupted, char)
					isCorrupted = true
					break
				}
			}
		}

		if !isCorrupted {
			incomplete = append(incomplete, runeSliceToString(openStack))
		}
	}

	return
}

func runeSliceToString(input []rune) string {
	var result string

	for _, rune := range input {
		result += string(rune)
	}

	return result
}

func pairOpener(char rune) rune {
	switch char {
	case ')':
		return '('
	case '}':
		return '{'
	case ']':
		return '['
	case '>':
		return '<'
	default:
		panic(char)
	}
}

func pairCloser(char rune) rune {
	switch char {
	case '(':
		return ')'
	case '{':
		return '}'
	case '[':
		return ']'
	case '<':
		return '>'
	default:
		panic(char)
	}
}
