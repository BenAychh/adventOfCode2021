package sonar

import (
	"adventOfCode2021/util"
	"fmt"
	"log"
)

func Depth() {
	dataStr, err := util.ReadFileIntoLines("day1_1.txt")
	if err != nil {
		log.Fatalf("failed to read file %v", err)
	}
	data, err := util.ToIntArray(dataStr)
	if err != nil {
		log.Fatalf("failed to convert to ints %v", err)
	}
	fmt.Println(CountDepthIncreases(data, 1))
	fmt.Println(CountDepthIncreases(data, 3))
}

func CountDepthIncreases(depths []int, groupsOf int) (count int) {
	if len(depths) == 0 {
		return
	}
	previous := 0
	for i := 0; i < groupsOf; i++ {
		previous += depths[i]
	}
	totalEntries := len(depths)
	for index := range depths[1:] {
		if index + groupsOf >= totalEntries {
			break
		}
		total := 0
		for i := 0; i < groupsOf; i++ {
			total += depths[index + i + 1]
		}
		if total > previous {
			count++
		}
		previous = total
	}
	return
}


