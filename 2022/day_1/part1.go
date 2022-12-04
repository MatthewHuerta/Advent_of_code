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
	largest := 0
	for _, v := range elves {
		if v > largest {
			largest = v
		}
	}
	fmt.Println(largest)
}
