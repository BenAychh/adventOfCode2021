package bingo

import (
	"regexp"
	"strconv"
	"strings"
)

type Game struct {
	draws  []int
	boards []*Board
}

func NewGame(data []string) *Game {
	drawsStr := strings.Split(data[0], ",")
	g := &Game{
		draws:  make([]int, len(drawsStr)),
		boards: make([]*Board, 0),
	}
	for index, str := range drawsStr {
		num, _ := strconv.Atoi(str)
		g.draws[index] = num
	}
	for i := 2; i < len(data); i += 6 {
		g.boards = append(g.boards, NewBoard(data[i:i+5]))
	}
	return g
}

func (g *Game) Run() (*Board, int) {
	for _, draw := range g.draws {
		for _, board := range g.boards {
			board.MarkNumber(draw)
			if board.DidWin() {
				return board, draw
			}
		}
	}
	return nil, 0
}

func (g *Game) Reset() {
	for _, b := range g.boards {
		b.alreadyWon = false
		for _, row := range b.data {
			for _, cell := range row {
				cell.Matched = false
			}
		}
	}
}

func (g *Game) RunUntilEveryBoardWins() (*Board, int) {
	wonBoardsCount := 0
	for _, draw := range g.draws {
		for _, board := range g.boards {
			board.MarkNumber(draw)
			if !board.alreadyWon && board.DidWin() {
				wonBoardsCount++
				board.alreadyWon = true
			}
			if wonBoardsCount == len(g.boards) {
				return board, draw
			}
		}
	}
	return nil, 0
}

type Square struct {
	Value int
	Matched bool
}

type Board struct {
	alreadyWon bool
	data [5][5]*Square
}

var anySpaces = regexp.MustCompile("\\s+")

func NewBoard(lines []string) *Board {
	b := &Board{data: [5][5]*Square{}}
	for index, line := range lines[:5] {
		split := anySpaces.Split(strings.TrimSpace(line), -1)
		for jndex, str := range split {
			square, _ := strconv.Atoi(str)
			b.data[index][jndex] = &Square{
				Value:   square,
			}
		}
	}
	return b
}

func (b *Board) MarkNumber(v int) {
	out:
	for _, row := range b.data {
		for _, cell := range row {
			if cell.Value == v {
				cell.Matched = true
				break out
			}
		}
	}
}

func (b *Board) DidWin() bool {
	for i := 0; i < 5; i++ {
		allMatch := true
		for j := 0; j < 5; j++ {
			if !b.data[i][j].Matched {
				allMatch = false
				break
			}
		}
		if allMatch {
			return true
		}
	}

	for i := 0; i < 5; i++ {
		allMatch := true
		for j := 0; j < 5; j++ {
			if !b.data[j][i].Matched {
				allMatch = false
				break
			}
		}
		if allMatch {
			return true
		}
	}
	return false
}

func (b *Board) UnmatchedSum() int {
	sum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !b.data[i][j].Matched {
				sum += b.data[i][j].Value
			}
		}
	}
	return sum
}
