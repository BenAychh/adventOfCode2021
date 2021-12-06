package util

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadFileIntoLines(name string) ([]string) {
	body, err := ioutil.ReadFile(fmt.Sprintf("./inputs/%s", name))
	if err != nil {
		panic("failed to read file")
	}
	return strings.Split(string(body), "\n")
}

func ToIntArray(arr []string) ([]int, error) {
	results := make([]int, len(arr))
	for index, item := range arr {
		num, err := strconv.Atoi(item)
		if err != nil {
			return nil, err
		}
		results[index] = num
	}
	return results, nil
}