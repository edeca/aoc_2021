// Advent of Code 2022 - Day 2
//
// Aim is to learn some golang - this is the second golang program I have written

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculate_position_from_file(name string) {
	var horizontal, depth, aim int = 0, 0, 0

	fh, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
	}
	defer fh.Close()

	scanner := bufio.NewScanner(fh)

	// Not needed as this is the default
	//scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		// Split on whitespace to get command and numeric argument
		fields := strings.Fields(scanner.Text())
		// Convert the numeric argument to an integer
		distance, _ := strconv.Atoi(fields[1])

		switch fields[0] {
		case "forward":
			horizontal += distance
			depth += aim * distance

		case "down":
			aim += distance

		case "up":
			aim -= distance

		default:
			panic("Invalid command found")
		}
	}

	// Depth in part 1 is equivalent to aim in part 2, so use it here
	fmt.Println("Part 1 - horizontal:", horizontal, ", depth:", aim, ", answer:", horizontal*aim)
	fmt.Println("Part 2 - horizontal:", horizontal, ", depth:", depth, ", answer:", horizontal*depth)
}

func main() {
	calculate_position_from_file("input.txt")
}
