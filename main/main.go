package main

import (
	"adventOfCode2021/hydrothermal"
	"adventOfCode2021/util"
	"fmt"
)

func main() {
	// day 1
	//sonar.Depth()

	// day 2
	//s := sub.Sub{}
	//data, err := util.ReadFileIntoLines("day2.txt")
	//if err != nil {
	//	log.Fatalf("failed to read file %v", err)
	//}
	//s.BasicMode = true
	//s.Move(data)
	//fmt.Printf("h: %d, v: %d, h * v: %d\n", s.HorizontalPos, s.VerticalPos, s.HorizontalPos * s.VerticalPos)
	//s.Reset()
	//s.Move(data)
	//fmt.Printf("h: %d, v: %d, h * v: %d\n", s.HorizontalPos, s.VerticalPos, s.HorizontalPos * s.VerticalPos)

	// Day 3
	//d := diagnosis.Diag{}
	//data, err := util.ReadFileIntoLines("day3.txt")
	//if err != nil {
	//	log.Fatalf("failed to read file %v", err)
	//}
	//d.ReadAll(data)
	////fmt.Printf("%d * %d = %d\n", d.Gamma(), d.Epsilon(), d.Gamma() * d.Epsilon())
	//fmt.Println(d.Oxygen(0), d.Scrubber(0))

	// Day 4
	//data, err := util.ReadFileIntoLines("day4.txt")
	//if err != nil {
	//	log.Fatalf("failed to read file %v", err)
	//}
	//game := bingo.NewGame(data)
	//board, lastNumber := game.Run()
	//fmt.Println(board.UnmatchedSum(), lastNumber)
	//game.Reset()
	//board, lastNumber = game.RunUntilEveryBoardWins()
	//fmt.Println(board.UnmatchedSum(), lastNumber)

	// Day 5
	data := util.ReadFileIntoLines("day5.txt")
	m := hydrothermal.Map{}
	m.AddLines(data)
	m.Plot()
	fmt.Println(m.PointsWithVentCountGreaterThan(1))
}
