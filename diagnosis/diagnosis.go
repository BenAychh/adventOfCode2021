package diagnosis

import (
	"strconv"
)

type Diag struct {
	Data []string
}

func (d *Diag) ReadAll(lines []string) {
	d.Data = lines
}

func (d *Diag) ReadLine(line string) {
	d.Data = append(d.Data, line)
}

func (d *Diag) Counts() ([]rune, []rune) {
	if len(d.Data) == 0 {
		return nil, nil
	}

	oneCount := make([]int, len(d.Data[0]))
	zeroCount := make([]int, len(d.Data[0]))

	for _, datum := range d.Data {
		for index, c := range datum {
			if c == '0' {
				zeroCount[index]++
			} else {
				oneCount[index]++
			}
		}
	}

	most := make([]rune, len(d.Data[0]))
	least := make([]rune, len(d.Data[0]))

	for index, count := range zeroCount {
		if oneCount[index] < count {
			most[index] = '0'
			least[index] = '1'
		} else {
			most[index] = '1'
			least[index] = '0'
		}
	}


	return most, least
}

func (d *Diag) Gamma() int64 {
	mosts, _ := d.Counts()
	res, _ := strconv.ParseInt(string(mosts), 2, 64)
	return res
}

func (d *Diag) Epsilon() int64 {
	_, leasts := d.Counts()
	res, _ := strconv.ParseInt(string(leasts), 2, 64)
	return res
}

func (d *Diag) Oxygen(index int) int64 {
	mosts, _ := d.Counts()
	c := mosts[index]
	newData := make([]string, 0)
	for _, datum := range d.Data {
		if rune(datum[index]) == c {
			newData = append(newData, datum)
		}
	}
	if len(newData) == 1{
		res, _ := strconv.ParseInt(newData[0], 2, 64)
		return res
	}
	d2 := Diag{Data: newData}
	return d2.Oxygen(index + 1)
}

func (d *Diag) Scrubber(index int) int64 {
	_, leasts := d.Counts()
	c := leasts[index]
	newData := make([]string, 0)
	for _, datum := range d.Data {
		if rune(datum[index]) == c {
			newData = append(newData, datum)
		}
	}
	if len(newData) == 1{
		res, _ := strconv.ParseInt(newData[0], 2, 64)
		return res
	}
	d2 := Diag{Data: newData}
	return d2.Scrubber(index + 1)
}