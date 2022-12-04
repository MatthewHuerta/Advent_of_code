package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	dat, _ := os.Open("input.txt")
	defer dat.Close()
	scanner := bufio.NewScanner(dat)
	var elves []int
	sum_per_elf := 0
	for scanner.Scan() {
		integer_val, _ := strconv.Atoi(scanner.Text())
		if scanner.Text() == "" {
			elves = append(elves, sum_per_elf+integer_val)
			sum_per_elf = 0
		} else {
			sum_per_elf += integer_val
		}
	}
	// for _, v := range elves {
	// 	if v >= 64499 {
	// 		fmt.Println(v)
	// 	}
	// }
	largest := 0
	second := 0
	third := 0
	for _, v := range elves {
		if v > largest {
			x := largest
			largest = v
			v = x
		}
		if v > second {
			x := second
			second = v
			v = x
		}
		if v > third {
			third = v
		}
	}
	fmt.Println(largest + second + third)
	fmt.Println(largest)
	fmt.Println(second)
	fmt.Println(third)
}
