package hydrothermal

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func NewPoint(str string) *Point {
	coordinates := strings.Split(str, ",")
	if len(coordinates) != 2 {
		panic("did not get exactly 2 points")
	}
	x, _ := strconv.Atoi(coordinates[0])
	y, _ := strconv.Atoi(coordinates[1])
	return &Point{x, y}
}

type Line struct {
	p1 *Point
	p2 *Point
}

func NewLine(str string) *Line {
	pointsStr := strings.Split(str, " -> ")
	if len(pointsStr) != 2 {
		panic("points string does not have exactly 2 points")
	}
	return &Line{p1: NewPoint(pointsStr[0]), p2: NewPoint(pointsStr[1])}
}

func (l *Line) IsVertical() bool {
	return l.p1.x == l.p2.x
}

func (l *Line) IsHorizontal() bool {
	return l.p1.y == l.p2.y
}

func (l *Line) Slope() float64 {
	return float64(l.p1.y - l.p2.y) / float64(l.p1.x - l.p2.x)
}

func (l *Line) Is45Degree() bool {
	return math.Abs(l.Slope()) == 1
}

func (l *Line) Points() (results []*Point) {
	points := l.OrderedCoordinates()
	if l.IsVertical() {
		for y := points[0].y; y <= points[1].y; y++ {
			results = append(results, &Point{points[0].x, y})
		}
	} else if l.IsHorizontal() {
		for x := points[0].x; x <= points[1].x; x++ {
			results = append(results, &Point{x, points[0].y})
		}
	} else if l.Is45Degree() {
		y := points[0].y
		yDelta := int(l.Slope())
		for x := points[0].x; x <= points[1].x; x++ {
			results = append(results, &Point{x, y})
			y += yDelta
		}
	}
	return results
}

func (l *Line) OrderedCoordinates() [2]*Point {
	if l.IsVertical() {
		if l.p1.y < l.p2.y {
			return [2]*Point{l.p1, l.p2}
		}
		return [2]*Point{l.p2, l.p1}
	} else if l.IsHorizontal() || l.Is45Degree(){
		if l.p1.x < l.p2.x {
			return [2]*Point{l.p1, l.p2}
		}
		return [2]*Point{l.p2, l.p1}
	}
	panic("unknown line")
}

type Map struct {
	VentCounts map[Point]int
	Lines []*Line
}

func (m *Map) AddLines(lines []string) {
	for _, line := range lines {
		m.AddLine(line)
	}
}

func (m *Map) AddLine(line string) {
	m.Lines = append(m.Lines, NewLine(line))
}

func (m *Map) Plot() {
	if m.VentCounts == nil {
		m.VentCounts = make(map[Point]int)
	}
	zero := Point{0, 0}
	for _, l1 := range m.Lines {
		points := l1.Points()
		for _, point := range points {
			if *point == zero {
				fmt.Println("here?")
			}
			m.VentCounts[*point]++
		}
	}
}

func (m *Map) PointsWithVentCountGreaterThan(count int) (result int) {
	for point, v := range m.VentCounts {
		if v > count {
			fmt.Println(point, v)
			result++
		}
	}
	return result
}