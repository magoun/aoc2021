package giantSquid

import (
	"aoc2021/helpers"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {

	dat, err := os.ReadFile("giantSquid/input")
	// dat, err := os.ReadFile("giantSquid/test")
	helpers.Check(err)

	stringDat := string(dat)
	stringArr := strings.Split(stringDat, "\n")

	playBingo(stringArr)
	loseBingo(stringArr)

}

func loseBingo(input []string) int {
	callOrder, boards := getCallOrderAndBingoBoards(input)
	lastWinningScore := 0

	for _, calledNumber := range callOrder {
		for _, board := range boards {
			if board.WinningScore <= 0 {
				winningScore := board.Mark(calledNumber)

				if winningScore > 0 {
					lastWinningScore = winningScore
				}
			}
		}
	}

	output := fmt.Sprintf("The losing board score is %d.", lastWinningScore)
	fmt.Println(output)
	return lastWinningScore
}

func playBingo(input []string) int {
	callOrder, boards := getCallOrderAndBingoBoards(input)

	for _, calledNumber := range callOrder {
		for _, board := range boards {
			winningScore := board.Mark(calledNumber)

			if winningScore > 0 {
				// board.Print()
				output := fmt.Sprintf("The winning board score is %d.", winningScore)
				fmt.Println(output)
				return winningScore
			}
		}
	}

	return -1
}

func getCallOrderAndBingoBoards(input []string) ([]string, []*BingoBoard) {
	callString := input[0]
	callOrder := strings.Split(callString, ",")

	line := 2
	var boards []*BingoBoard

	for line < len(input) {
		boardSlice := input[line : line+5]
		bb := newBingoBoard(boardSlice)
		boards = append(boards, bb)
		line += 6
	}

	return callOrder, boards
}

type BingoBoard struct {
	BoardRows    []string
	Rows         [5][5]string
	Columns      [5][5]string
	Diagonals    [2][5]string
	WinningScore int
}

func newBingoBoard(rawBoard []string) *BingoBoard {
	bb := BingoBoard{BoardRows: rawBoard}

	for row, rowString := range bb.BoardRows {

		rowArr := strings.Fields(rowString)

		for col, value := range rowArr {
			bb.Rows[row][col] = value
			bb.Columns[col][row] = value

			if row == col {
				bb.Diagonals[0][row] = value
			}

			if row+col == 4 {
				bb.Diagonals[1][row] = value
			}
		}
	}

	return &bb
}

func (bb *BingoBoard) Mark(number string) int {
	bb.CheckRows(number)
	bb.CheckColumns(number)
	// bb.CheckDiagonals(number)

	return bb.WinningScore
}

func (bb *BingoBoard) CheckRows(number string) {
	for line, values := range bb.Rows {
		complete := true

		for place, value := range values {
			if value == number {
				bb.Rows[line][place] = "x"
			} else if value != "x" {
				complete = false
			}
		}

		if complete {
			bb.CalculateBoardScore(number)
		}
	}
}

func (bb *BingoBoard) CheckColumns(number string) {
	for line, values := range bb.Columns {
		complete := true

		for place, value := range values {
			if value == number {
				bb.Columns[line][place] = "x"
			} else if value != "x" {
				complete = false
			}
		}

		if complete {
			bb.CalculateBoardScore(number)
		}
	}
}

func (bb *BingoBoard) CheckDiagonals(number string) {
	for line, values := range bb.Diagonals {
		complete := true

		for place, value := range values {
			if value == number {
				bb.Diagonals[line][place] = "x"
			} else if value != "x" {
				complete = false
			}
		}

		if complete {
			bb.CalculateBoardScore(number)
		}
	}
}

func (bb *BingoBoard) CalculateBoardScore(lastValue string) {
	score := 0
	lastValueInt, _ := strconv.Atoi(lastValue)

	for _, values := range bb.Rows {
		for _, value := range values {
			if value != "x" {
				intValue, _ := strconv.Atoi(value)
				score += intValue
			}
		}
	}

	bb.WinningScore = score * lastValueInt
}

func (bb *BingoBoard) Print() {
	fmt.Println(bb.Rows)
	fmt.Println(bb.Columns)
	fmt.Println(bb.Diagonals)
}
