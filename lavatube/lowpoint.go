package lavatube

import (
	"fmt"
	"sort"
	"strconv"
)

type LavaMap struct {
	data [][]int
}

func (l *LavaMap) AddLines(lines []string) {
	for _, line := range lines {
		l.AddLine(line)
	}
}

func (l *LavaMap) AddLine(str string) {
	intLine := make([]int, len(str))
	for index, r := range str {
		num, _ := strconv.Atoi(string(r))
		intLine[index] = num
	}
	l.data = append(l.data, intLine)
}

func (l *LavaMap) LowPoints() [][2]int {
	resp := make([][2]int, 0)
	for y := range l.data {
		for x := range l.data[y] {
			val, _ := l.ValueAt(x, y)
			neighbors := l.getNeighbors(x, y)
			isLowpoint := true
			for _, n := range neighbors {
				nValue, _ := l.ValueAt(n[0], n[1])
				if val >= nValue {
					isLowpoint = false
					break
				}
			}
			if isLowpoint {
				resp = append(resp, [2]int{x, y})
			}
		}
	}
	return resp
}

func (l *LavaMap) ValuesAt(coords [][2]int) (resp []int) {
	for _, c := range coords {
		if v, ok := l.ValueAt(c[0], c[1]); ok {
			resp = append(resp, v)
		}
	}
	return
}

type CoordSet struct {
	coords [][2]int
}

func NewCoordSet() *CoordSet {
	return &CoordSet{coords: [][2]int{}}
}

func (c *CoordSet) Add(x, y int) {
	c.coords = append(c.coords, [2]int{x, y})
}

func (c *CoordSet) InSet(coord [2]int) bool {
	for _, c := range c.coords {
		if c[0] == coord[0] && c[1] == coord[1]{
			return true
		}
	}
	return false
}

func (l *LavaMap) BasinSizes(points [][2]int) []int {
	resp := make([]int, len(points))
	for index, p := range points {
		resp[index] = l.BasinSize(p[0], p[1], NewCoordSet())
	}
	sort.Slice(resp, func(i, j int) bool {
		return resp[j] < resp[i]
	})
	return resp
}

func (l *LavaMap) BasinSize(x, y int, touchedCoords *CoordSet) int {
	_, ok := l.ValueAt(x, y)
	if !ok {
		return 0
	}
	if x == 9 && y == 2 {
		fmt.Println(x, y)
	}
	touchedCoords.Add(x, y)
	neighbors := l.getNeighbors(x, y)
	nValues := l.ValuesAt(neighbors)
	filteredNeighbors := make([][2]int, 0)
	for index, n := range neighbors {
		if touchedCoords.InSet(n) {
			continue
		}
		if nValues[index] == 9 {
			continue
		}
		filteredNeighbors = append(filteredNeighbors, n)
		touchedCoords.Add(n[0], n[1])
	}
	count := 0
	for _, fn := range filteredNeighbors {
		count += l.BasinSize(fn[0], fn[1], touchedCoords)
	}
	return 1 + count
}

func RiskLevels(vals []int) (resp int) {
	for _, v := range vals {
		resp += v + 1
	}
	return
}

func (l *LavaMap) ValueAt(x, y int) (int, bool) {
	if y < 0 || y >= len(l.data) {
		return 0, false
	}
	if x < 0 || x >= len(l.data[y]) {
		return 0, false
	}
	return l.data[y][x], true
}

func (l *LavaMap) getNeighbors(x int, y int) (resp [][2]int) {
	xl := x - 1
	xr := x + 1
	ya := y - 1
	yb := y + 1
	if _, ok := l.ValueAt(xl, y); ok {
		resp = append(resp, [2]int{xl, y})
	}
	if _, ok := l.ValueAt(xr, y); ok {
		resp = append(resp, [2]int{xr, y})
	}
	if _, ok := l.ValueAt(x, ya); ok {
		resp = append(resp, [2]int{x, ya})
	}
	if _, ok := l.ValueAt(x, yb); ok {
		resp = append(resp, [2]int{x, yb})
	}
	return
}
