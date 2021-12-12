package sevenseg

import (
	"fmt"
	"sort"
	"strings"
)

func UniqueCount(data []string) {
	sb := strings.Builder{}
	for _, d := range data {
		split := strings.Split(d, " | ")
		sb.WriteString(" " + split[1])
	}
	split := strings.Split(sb.String(), " ")
	count := 0
	for _, s := range split {
		switch len(s) {
		case 2, 3, 4, 7:
			count++
		}
	}
	fmt.Println(count)
}

func SumOutput(data []string) {
	total := 0
	for _, d := range data {
		total += Decode(d)
	}
	fmt.Println(total)
}

func Decode(data string) int {
	split := strings.Split(data, " | ")
	start := strings.Split(split[0], " ")
	output := strings.Split(split[1], " ")
	allNumbers := append(start, output...)
	var perms = permutations("dabcgef")
	var order string
	for _, p := range perms {
		works := true
		for _, n := range allNumbers {
			num := getNumber(p, n)
			if num == -1 {
				works = false
				break
			}
		}
		if works {
			order = p
			break
		}
	}

	if order == "" {
		panic(fmt.Sprintf("failed to find solution for %d", data))
	}

	numbers := make([]int, len(output))
	for index, o := range output {
		numbers[index] = getNumber(order, o)
	}

	return numbers[0] * 1000 + numbers[1] * 100 + numbers[2] * 10 + numbers[3]
}

var numbers = [][]int{
	{1, 2, 3, 4, 5, 6},    // 0
	{2, 3},                // 1
	{1, 2, 4, 5, 7},       // 2
	{1, 2, 3, 4, 7},       // 3
	{2, 3, 6, 7},          // 4
	{1, 3, 4, 6, 7},       // 5
	{1, 3, 4, 5, 6, 7},    // 6
	{1, 2, 3},             // 7
	{1, 2, 3, 4, 5, 6, 7}, // 8
	{1, 2, 3, 4, 6, 7},    // 9
}

func getNumber(order string, nStr string) int {
	m := positionMap(order)
	n := getOrderedNumber(m, nStr)
	for i, number := range numbers {
		if len(n) != len(number) {
			continue
		}
		match := true
		for index := range n {
			if n[index] != number[index] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}
	return -1
}

func positionMap(order string) map[rune]int {
	resp := make(map[rune]int)
	for index, r := range order {
		resp[r] = index + 1
	}
	return resp
}

func getOrderedNumber(m map[rune]int, nStr string) []int {
	resp := make([]int, len(nStr))
	for index, r := range nStr {
		resp[index] = m[r]
	}
	sort.Slice(resp, func(i, j int) bool {
		return resp[i] < resp[j]
	})
	return resp
}

func join(ins []rune, c rune) (result []string) {
	for i := 0; i <= len(ins); i++ {
		result = append(result, string(ins[:i])+string(c)+string(ins[i:]))
	}
	return
}

func permutations(testStr string) []string {
	var n func(testStr []rune, p []string) []string
	n = func(testStr []rune, p []string) []string {
		if len(testStr) == 0 {
			return p
		} else {
			result := []string{}
			for _, e := range p {
				result = append(result, join([]rune(e), testStr[0])...)
			}
			return n(testStr[1:], result)
		}
	}

	output := []rune(testStr)
	return n(output[1:], []string{string(output[0])})
}
