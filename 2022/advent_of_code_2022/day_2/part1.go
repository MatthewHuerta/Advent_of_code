package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	sign  string
	value int
	Prev  *Node
	Next  *Node
}

func (this *Node) vs(other *Node) int {
	// draw case
	if this == other {
		return (3 + this.value)
	}
	// loss case
	if this.Next == other {
		return (0 + this.value)
	}
	// win case
	if this.Prev == other {
		return (6 + this.value)
	}
	return 0
}

func main() {
	var rock = Node{value: 1, sign: "rock"}
	var paper = Node{value: 2, sign: "paper"}
	var scissors = Node{value: 3, sign: "scissors"}
	rock.Prev = &scissors
	rock.Next = &paper
	paper.Prev = &rock
	paper.Next = &scissors
	scissors.Prev = &paper
	scissors.Next = &rock

	input, _ := os.Open("input.txt")
	defer input.Close()
	scanner := bufio.NewScanner(input)
	score := 0
	var elf *Node
	var you *Node
	for scanner.Scan() {
		switch {
		case strings.Split(scanner.Text(), " ")[0] == "A":
			elf = &rock
		case strings.Split(scanner.Text(), " ")[0] == "B":
			elf = &paper
		case strings.Split(scanner.Text(), " ")[0] == "C":
			elf = &scissors
		}
		switch {
		case strings.Split(scanner.Text(), " ")[1] == "X":
			you = &rock
		case strings.Split(scanner.Text(), " ")[1] == "Y":
			you = &paper
		case strings.Split(scanner.Text(), " ")[1] == "Z":
			you = &scissors
		}
		this_round := you.vs(elf)
		score += this_round

	}
	fmt.Println(score)
}
