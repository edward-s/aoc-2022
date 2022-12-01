package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	var elves []int
	var sum = 0

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		if n, err := strconv.Atoi(scanner.Text()); err == nil {
			sum += n
		} else {
			elves = append(elves, sum)
			sum = 0
		}
	}

	sort.Ints(elves)
	lastThree := elves[len(elves)-3:]

	var sumLastThree = 0
	for _, n := range lastThree {
		sumLastThree += n
	}

	fmt.Println(lastThree)
	fmt.Println(sumLastThree)
}
