package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Stack struct {
	top_crate *string
	column    []string
	height    int
}

// type Columns struct {
// 	stacks map[string]*Stack
// }

// func (c *Columns) AddCrate(stack string, crates ...rune) {

// 	for _, crate := range crates {
// 		stack.column = append(stack.column, crate)
// 		stack.height += 1
// 		stack.top_crate = &stack.column[len(stack.column)-1]
// 	}
// }

func NewStack(crates ...string) *Stack {
	var stack Stack
	stack.height = 0
	stack.top_crate = nil
	for _, crate := range crates {
		stack.AddCrate(crate)
	}
	return &stack
}

func (stack *Stack) AddCrate(crates ...string) {
	for _, crate := range crates {
		stack.column = append(stack.column, crate)
		stack.height += 1
		stack.top_crate = &stack.column[len(stack.column)-1]
	}
}

func (stack *Stack) RemoveCrate() {
	stack.column = stack.column[:len(stack.column)-1]
	stack.height -= 1
	if stack.height <= 0 {
		stack.top_crate = nil
	} else {
		stack.top_crate = &stack.column[len(stack.column)-1]
	}
}

func (source *Stack) MoveCrate(target *Stack) {
	target.AddCrate(*source.top_crate)
	source.RemoveCrate()
}

func (source *Stack) MoveCrates(numberCrates int, target *Stack) {
	var holderStack Stack
	holderStack.top_crate = nil
	for i := 1; i <= numberCrates; i++ {
		source.MoveCrate(&holderStack)
	}
	for i := 1; i <= numberCrates; i++ {
		holderStack.MoveCrate(target)
	}

}

func main() {
	dat, _ := os.Open("input.txt")
	defer dat.Close()
	scanner := bufio.NewScanner(dat)
	number, _ := regexp.Compile("([0-9])")
	var lines = []string{}
	for scanner.Scan() {
		match := number.MatchString(scanner.Text())
		lines = append([]string{scanner.Text()}, lines...)
		if match {
			break
		}
	}
	column := make(map[int]string)
	stacks := make(map[string]*Stack)
	rowNumbers := number.FindAllStringSubmatchIndex(lines[0], 10)
	for _, r := range rowNumbers {
		num := string(lines[0][r[2]])
		column[r[2]] = num
		stacks[column[r[2]]] = NewStack()
		// fmt.Printf("%s:%d ", num, r[2])
		// fmt.Println(column[r[2]])
	}
	crate, _ := regexp.Compile("\\[(?P<c>[A-Z])\\]")
	for v := 1; v < len(lines); v++ {
		crates := crate.FindAllStringSubmatchIndex(lines[v], 10)
		for _, c := range crates {
			// fmt.Printf("%d:%c ", c[2], lines[v][c[2]])
			stacks[column[c[2]]].AddCrate(string(lines[v][c[2]]))
			// fmt.Println(column[c[2]])
		}
		// fmt.Println("")
	}
	number, _ = regexp.Compile("\\d+")
	scanner.Scan()
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		matches := number.FindAllString(scanner.Text(), 3)
		ntimes, _ := strconv.Atoi(matches[0])
		source := matches[1]
		target := matches[2]
		// fmt.Printf("n:%d s:%d t:%d", ntimes, source, target)
		// fmt.Println("")
		for i := 1; i <= ntimes; i++ {
			stacks[source].MoveCrate(stacks[target])
		}
	}
	for i := 1; i < 10; i++ {
		I := strconv.Itoa(i)
		fmt.Println(stacks[I].column)
	}
	for i := 1; i < 10; i++ {
		I := strconv.Itoa(i)
		fmt.Printf("%s", *stacks[I].top_crate)
	}
	fmt.Println("")

}
