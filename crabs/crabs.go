package crabs

import (
	"math"
	"strconv"
	"strings"
)

func Crabs(input string) int {
	split := strings.Split(input, ",")
	data := make([]int, len(split))
	for index, str := range split {
		data[index], _ = strconv.Atoi(str)
	}
	min := data[0]
	max := data[0]
	for _, datum := range data {
		if datum < min {
			min = datum
		}
		if datum > max {
			max = datum
		}
	}
	totalFuel := math.MaxInt
	for i := min; i <= max; i++ {
		fuel := 0
		for _, datum := range data {
			fuel += int(math.Abs(float64(i) - float64(datum)))
		}
		if fuel < totalFuel {
			totalFuel = fuel
		}
	}
	return totalFuel
}

func Crabs2(input string) int {
	split := strings.Split(input, ",")
	data := make([]int, len(split))
	for index, str := range split {
		data[index], _ = strconv.Atoi(str)
	}
	min := data[0]
	max := data[0]
	for _, datum := range data {
		if datum < min {
			min = datum
		}
		if datum > max {
			max = datum
		}
	}
	totalFuel := math.MaxInt
	for i := min; i <= max; i++ {
		fuel := 0
		for _, datum := range data {
			n := int(math.Abs(float64(i) - float64(datum)))
			fuel += n * (n + 1) / 2
		}
		if fuel < totalFuel {
			totalFuel = fuel
		}
	}
	return totalFuel
}
