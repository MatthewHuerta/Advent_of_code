package main

import (
	"bufio"
	"fmt"
	"os"
)

func containsDuplicates(section string) bool {
	for c := 0; c < len(section); c++ {
		for m := 0; m < len(section); m++ {
			if m == c {
				continue
			}
			if section[m] == section[c] {
				return true
			}
		}
	}
	return false
}

func main() {
	dat, _ := os.Open("input.txt")
	defer dat.Close()
	scanner := bufio.NewScanner(dat)
	for scanner.Scan() {
		cipher := scanner.Text()
		for i := 0; i < len(cipher)-13; i++ {
			section := cipher[i : i+14]
			if !containsDuplicates(section) {
				fmt.Println(section, "at:", i+14)
				break
			} // else {
			// 	// fmt.Println(section, "contains no duplicates")
			// }
		}

	}

}
