// Advent of Code 2022 - Day 1
//
// Aim is to learn some golang - this is the first golang program I have written

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func count_higher_numbers_from_file(name string) int {
	// The first line will always be higher than 0, so start counting at -1
	higher := -1

	// Alternatively we could set prev to the max int by importing math
	// prev := math.MaxInt
	prev := 0

	fh, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
	}
	defer fh.Close()

	scanner := bufio.NewScanner(fh)

	// Not needed as this is the default
	//scanner.Split(bufio.ScanLines)

	// This returns slightly the wrong result because the first
	// line is higher than zero. We adjust for this above, but the
	// answer would be wrong for a file with no higher values.
	for scanner.Scan() {
		cur, _ := strconv.Atoi(scanner.Text())
		// TODO: What to do with err
		if cur > prev {
			higher += 1
		}
		prev = cur
	}

	return higher
}

// This approach is designed to stream the data, rather than read it all into memory
// at once. This means we only ever keep 4x ints in memory, plus the counter.
//
// Totally unnecessary for this AoC challenge which would fit into a tiny amount of
// RAM, but interesting for other use-cases.
func count_higher_windows_from_file(name string) int {
	var higher = 0
	var window = [4]int{0, 0, 0, 0}

	fh, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
	}
	defer fh.Close()

	scanner := bufio.NewScanner(fh)

	// Preload the first three values. The main loop will shift these and load the next,
	// so we only care about loading the first window.
	for n := 1; n < 4; n++ {
		scanner.Scan()
		window[n], err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println(window)

	// Keep a sliding window over 4 values, add first three and second three to compare
	for scanner.Scan() {
		cur, _ := strconv.Atoi(scanner.Text())

		// Shift all current values down
		for n := 0; n < 3; n++ {
			window[n] = window[n+1]
		}
		// Load in the newest value
		window[3] = cur

		sum1 := window[0] + window[1] + window[2]
		sum2 := window[1] + window[2] + window[3]
		//fmt.Println("sum1:", sum1, "sum2", sum2)

		if sum1 < sum2 {
			higher += 1
		}
	}

	return higher
}

func main() {
	var higher int

	higher = count_higher_numbers_from_file("input.txt")
	fmt.Println("There are", higher, "higher values")

	higher = count_higher_windows_from_file("input.txt")
	fmt.Println("There are", higher, "higher windows")

	// Curiously this does not cause an error 🤔
	// go seems to be quite permissive - unless we check the result
	// from the scanner object it silently returns
	higher = count_higher_windows_from_file("error_input.txt")
	fmt.Println("There are", higher, "higher windows")

	// This also doesn't crash the program. What a novelty.
	higher = count_higher_windows_from_file("i_dont_exist.txt")
	fmt.Println("There are", higher, "higher windows")
}
