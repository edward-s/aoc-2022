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
	fmt.Println(elves[len(elves)-1])
}
