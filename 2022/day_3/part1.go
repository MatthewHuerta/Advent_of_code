package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"unicode"
)

func to_value(r rune) int {
	if unicode.IsUpper(r) {
		return int(r) - 38
	}
	return int(r) - 96
}

func find_common(rucksack1 []rune, rucksack2 []rune) (rune, error) {
	var r rune
	for _, v := range rucksack1 {
		for _, c := range rucksack2 {
			if c == v {
				// fmt.Printf("%c : %d\n", v, to_value(v))
				return v, nil
			}
		}
	}
	err := errors.New("no match found")
	return r, err
}

func main() {
	dat, _ := os.Open("input.txt")
	defer dat.Close()
	scanner := bufio.NewScanner(dat)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		length := len(line)
		rucksack1 := line[0 : length/2]
		rucksack2 := line[length/2 : length]
		// fmt.Printf("%s(%d)\n", line, length)
		// fmt.Printf("%s(%d) : %s(%d)\n", rucksack1, len(rucksack1), rucksack2, len(rucksack2))
		common, err := find_common([]rune(rucksack1), []rune(rucksack2))
		if err != nil {
			log.Fatal(err)
		}
		total += to_value(common)

	}
	fmt.Println(total)
}
