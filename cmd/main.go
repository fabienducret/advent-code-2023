package main

import (
	"advent/day1"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("day1/data.txt")
	if err != nil {
		panic(err)
	}

	total := day1.GetTotalCalibrations(file)

	fmt.Println(total)
}
