package day1

import (
	"bufio"
	"io"
	"sort"
	"strconv"
	"strings"
)

var calibrationsMapping = []struct {
	key   string
	value int
}{
	{key: "1", value: 1},
	{key: "one", value: 1},
	{key: "2", value: 2},
	{key: "two", value: 2},
	{key: "3", value: 3},
	{key: "three", value: 3},
	{key: "4", value: 4},
	{key: "four", value: 4},
	{key: "5", value: 5},
	{key: "five", value: 5},
	{key: "6", value: 6},
	{key: "six", value: 6},
	{key: "7", value: 7},
	{key: "seven", value: 7},
	{key: "8", value: 8},
	{key: "eight", value: 8},
	{key: "9", value: 9},
	{key: "nine", value: 9},
}

type foundCalibration struct {
	index int
	value int
}

func GetTotalCalibrations(reader io.Reader) int {
	s := bufio.NewScanner(reader)
	var total int

	for s.Scan() {
		line := s.Text()
		var first, last int

		c := foundCalibrationsFrom(line)
		sortByIndex(c)
		first = firstOf(c)
		last = lastOf(c)

		lineTotal := concat(first, last)
		total += lineTotal
	}

	return total
}

func foundCalibrationsFrom(line string) []foundCalibration {
	var c []foundCalibration
	var firstIndex, lastIndex int

	for _, s := range calibrationsMapping {
		firstIndex = strings.Index(line, s.key)
		lastIndex = strings.LastIndex(line, s.key)

		if firstIndex != -1 {
			c = append(c, foundCalibration{firstIndex, s.value})
		}

		if lastIndex != -1 && firstIndex != lastIndex {
			c = append(c, foundCalibration{lastIndex, s.value})
		}
	}

	return c
}

func firstOf(c []foundCalibration) int {
	return c[0].value
}

func lastOf(c []foundCalibration) int {
	return c[len(c)-1].value
}

func concat(first, last int) int {
	concat := strconv.Itoa(first) + strconv.Itoa(last)
	t, _ := strconv.Atoi(concat)

	return t
}

func sortByIndex(c []foundCalibration) {
	sort.Slice(c, func(i, j int) bool {
		return c[i].index < c[j].index
	})
}
