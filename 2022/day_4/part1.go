package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func to_sections(elf string) [2]int {
	var sec [2]int
	sec[0], _ = strconv.Atoi(strings.Split(elf, "-")[0])
	sec[1], _ = strconv.Atoi(strings.Split(elf, "-")[1])
	return sec

}

// checks if elf1 contains elf2
func contains(elf1 [2]int, elf2 [2]int) bool {
	return elf1[0] <= elf2[0] && elf1[1] >= elf2[1]
}

func main() {
	dat, _ := os.Open("input.txt")
	defer dat.Close()
	scanner := bufio.NewScanner(dat)
	total := 0
	for scanner.Scan() {
		elf1 := to_sections(strings.Split(scanner.Text(), ",")[0])
		elf2 := to_sections(strings.Split(scanner.Text(), ",")[1])
		// fmt.Println(elf1, " contains ", elf2, " ? : ", contains(elf1, elf2))
		// fmt.Println(elf2, " contains ", elf1, " ? : ", contains(elf2, elf1))
		if contains(elf1, elf2) || contains(elf2, elf1) {
			total += 1
		}
	}
	fmt.Println(total)
}
