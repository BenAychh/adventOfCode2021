package main

import (
	"adventOfCode2021/chunk"
	"adventOfCode2021/util"
	"fmt"
	"sort"
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
	//start := time.Now()
	//data := util.ReadFileIntoLines("day5.txt")
	//m := hydrothermal.Map{}
	//m.AddLines(data)
	//m.Plot()
	//fmt.Println(m.PointsWithVentCountGreaterThan(1))
	//fmt.Println(time.Now().Sub(start))

	//start := time.Now()
	//data := util.ReadFileIntoLines("day6.txt")
	//laternfish.SetupOcean2(data[0])
	//for i := 0; i < 256; i++ {
	//	laternfish.TickOcean2()
	//}
	//fmt.Println(laternfish.CountOcean2())
	//
	//data := util.ReadFileIntoLines("day7.txt")
	//fmt.Println(crabs.Crabs2(data[0]))

	//data := util.ReadFileIntoLines("day8.txt")
	//sevenseg.SumOutput(data)

	//data := util.ReadFileIntoLines("day9.txt")
	//lm := lavatube.LavaMap{}
	//lm.AddLines(data)
	//lowPoints := lm.LowPoints()
	//fmt.Println(lm.BasinSize(9, 0, lavatube.NewCoordSet()))
	//fmt.Println(lm.BasinSizes(lowPoints))

	data := util.ReadFileIntoLines("day10.txt")
	incompleteLines := make([]string, 0)
	for _, d := range data {
		ok, _ := chunk.ProcessLine(d)
		if ok {
			incompleteLines = append(incompleteLines, d)
		}
	}
	incompleteScores := make([]int, len(incompleteLines))
	for index, line := range incompleteLines {
		c := chunk.Complete(line)
		incompleteScores[index] = chunk.ScoreCloses(c)
	}
	sort.Slice(incompleteScores, func(i, j int) bool {
		return incompleteScores[i] < incompleteScores[j]
	})
	fmt.Println(incompleteScores[len(incompleteScores) / 2])
}
