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

func find_common(rucksack1 []rune, rucksack2 []rune, rucksack3 []rune) (rune, error) {
	var r rune
	for _, v := range rucksack1 {
		for _, c := range rucksack2 {
			if c == v {
				for _, t := range rucksack3 {
					if c == t {

						// fmt.Printf("%c : %d\n", v, to_value(v))
						return v, nil
					}
				}
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
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		// fmt.Printf("processing line: %s\n", scanner.Text())
		if len(lines) == 3 {
			common, err := find_common([]rune(lines[0]), []rune(lines[1]), []rune(lines[2]))
			if err != nil {
				log.Fatal(err)
			}
			// fmt.Printf("(1)%s\n(2)%s\n(3)%s\ncommon: (%c)\n", lines[0], lines[1], lines[2], common)
			total += to_value(common)
			lines = nil
		}
	}
	fmt.Println(total)
}
