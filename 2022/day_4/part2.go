package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Section struct {
	beg int
	end int
}

func to_sections(secString string) Section {
	var sec Section
	sec.beg, _ = strconv.Atoi(strings.Split(secString, "-")[0])
	sec.end, _ = strconv.Atoi(strings.Split(secString, "-")[1])
	return sec
}

func (elf1 Section) contains(elf2 Section) bool {
	return elf1.beg <= elf2.beg && elf1.end >= elf2.end
}

// checks if elf1 and elf2 overlap
func (elf1 Section) overlaps(elf2 Section) bool {
	// case one:
	if elf1.beg <= elf2.beg && elf2.beg <= elf1.end {
		return true
		// case two:
	} else if elf1.beg <= elf2.end && elf2.end <= elf1.end {
		return true
	}
	return false
}

func main() {
	dat, _ := os.Open("input.txt")
	defer dat.Close()
	scanner := bufio.NewScanner(dat)
	total := 0
	for scanner.Scan() {
		elf1 := to_sections(strings.Split(scanner.Text(), ",")[0])
		elf2 := to_sections(strings.Split(scanner.Text(), ",")[1])
		// fmt.Println(elf1, " overlaps ", elf2, " ? : ", elf1.overlaps(elf2))
		if elf1.overlaps(elf2) || elf1.contains(elf2) || elf2.contains(elf1) {
			total += 1
		}
	}
	fmt.Println(total)
}
