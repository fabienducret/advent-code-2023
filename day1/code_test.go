package day1_test

import (
	"advent/day1"
	"strings"
	"testing"
)

func TestDay1(t *testing.T) {
	t.Run("get calibrations total for 1 line", func(t *testing.T) {
		line := "kspfzvvvfkztcs9threefoureightsixseveneight"
		reader := strings.NewReader(line)

		total := day1.GetTotalCalibrations(reader)

		assertEqual(t, total, 98)
	})

	t.Run("get calibrations total for multiple lines", func(t *testing.T) {
		line := "two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen"
		reader := strings.NewReader(line)

		total := day1.GetTotalCalibrations(reader)

		assertEqual(t, total, 281)
	})
}

func assertEqual(t *testing.T, got, want int) {
	if got != want {
		t.Errorf("error got %d, want %d", got, want)
	}
}
