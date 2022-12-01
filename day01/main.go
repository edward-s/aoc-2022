package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/edward-s/aoc-2022/utils"
)

func calculateCalories(lines []string) []int {
	var calories []int
	var sum = 0
	for _, line := range lines {
		if n, err := strconv.Atoi(line); err == nil {
			sum += n
		} else {
			calories = append(calories, sum)
			sum = 0
		}
	}

	sort.Ints(calories)
	return calories
}

func part1(lines []string) {
	calories := calculateCalories(lines)
	fmt.Println(calories[len(calories)-1])
}

func part2(lines []string) {
	calories := calculateCalories(lines)
	lastThree := calories[len(calories)-3:]

	var sumLastThree = 0
	for _, n := range lastThree {
		sumLastThree += n
	}
	fmt.Println(sumLastThree)
}

func main() {
	lines, _ := utils.ReadLines("input.txt")

	part1(lines)
	part2(lines)
}
