package laternfish

import (
	"strconv"
	"strings"
)

var fishCount = 0

type Reproducer interface {
	Tick()
	FamilyCount() int
}

type Ocean struct {
	Fish []*Fish
}

var Ocean2 = [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}

func NewOcean(str string) *Ocean {
	split := strings.Split(str, ",")
	o := &Ocean {
		Fish: make([]*Fish, len(split)),
	}
	for index, intStr := range split {
		timer, _ := strconv.Atoi(intStr)
		o.Fish[index] = NewFish(timer)
	}

	return o
}

func SetupOcean2(str string) {
	split := strings.Split(str, ",")
	for _, intStr := range split {
		timer, _ := strconv.Atoi(intStr)
		Ocean2[timer]++
	}
}

func TickOcean2() {
	spawnCount := Ocean2[0]

	for index, count := range Ocean2[1:] {
		Ocean2[index] = count
	}
	Ocean2[6] += spawnCount
	Ocean2[8] = spawnCount
}

func CountOcean2() int {
	count := 0
	for _, c := range Ocean2 {
		count += c
	}
	return count
}

func (o *Ocean) Tick(count int) {
	for i := 0; i < count; i++ {
		for _, fish := range o.Fish {
			fish.Tick()
		}
	}
}

func (o *Ocean) FishCount() int {
	count := 0
	for _, fish := range o.Fish {
		count += fish.FamilyCount()
	}
	return count
}

type Fish struct {
	Timer int
	ID int
	Children []Reproducer
}

func NewFish(timer int) *Fish {
	fishCount++
	return &Fish{
		ID:    fishCount,
		Timer: timer,
	}
}

func (f *Fish) Tick() {
	for _, c := range f.Children {
		c.Tick()
	}
	f.Timer--
	if f.Timer == -1 {
		f.Timer = 6
		f.Children = append(f.Children, NewFish(8))
	}
}

func (f *Fish) FamilyCount() int {
	count := 0
	for _, c := range f.Children {
		count += c.FamilyCount()
	}
	return count + 1
}