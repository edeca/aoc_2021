// Advent of Code 2022 - Day 3
//
// Aim is to learn some golang - this is the third golang program I have written

package main

import (
	"bufio"
	"fmt"
	"os"
)

// Sample file data:
//
// 011110010100
func calculate_diagnostics_from_file(name string) {
	// Hard coded to the length of the input string. Can be longer to support
	// a wider range of values.
	var ones [12]int
	var zeros [12]int

	fh, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
	}
	defer fh.Close()

	scanner := bufio.NewScanner(fh)

	for scanner.Scan() {
		line := scanner.Text()

		// Loop through the line using slicing, comparing the value to 0
		// There must be a more efficient way to do this using bitwise comparisons on the
		// integer value, rather than slicing strings?
		for i := 0; i < len(line); i++ {
			if line[i] == 0x30 {
				zeros[i] += 1
			} else {
				ones[i] += 1
			}
		}
	}

	// Now calculate the two values for gamma and epsilon, by comparing the
	// number of ones vs. zeros
	gamma := 0
	epsilon := 0
	for i := 0; i < len(ones); i++ {
		if ones[i] < zeros[i] {
			epsilon |= 0x800 >> i
		}
	}
	// gamma is the inverse, so use bitwise magic
	gamma = epsilon ^ 0xFFF

	fmt.Println("Part 1: gamma is", gamma, "and epsilon is", epsilon, "so the answer is", gamma*epsilon)
}

func main() {
	calculate_diagnostics_from_file("input.txt")
}
