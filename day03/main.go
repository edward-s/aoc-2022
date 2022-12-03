package main

import (
	"fmt"

	"github.com/edward-s/aoc-2022/utils"
)

const GROUP_SIZE = 3

func score(c rune) rune {
	if c >= 'A' && c <= 'Z' {
		return c - 'A' + 27
	} else {
		return c - 'a' + 1
	}
}

func part1(lines []string) rune {
	var sum rune = 0
	for _, line := range lines {
		compartment1 := line[:len(line)/2]
		compartment2 := line[len(line)/2:]
		chars := map[rune]bool{}
		for _, c := range compartment1 {
			chars[c] = true
		}

		found := map[rune]bool{}

		for _, c := range compartment2 {
			if chars[c] && !found[c] {
				found[c] = true
				sum += score(c)
			}
		}
	}

	return sum
}

func findCommonCharOfLines(lines []string) rune {
	commonChars := map[rune]int{}
	for _, line := range lines {
		found := map[rune]bool{}
		for _, c := range line {
			if !found[c] {
				found[c] = true
				if _, ok := commonChars[c]; ok {
					commonChars[c] = commonChars[c] + 1
				} else {
					commonChars[c] = 1
				}
			}
		}
	}

	var commonChar rune = 0
	for k, v := range commonChars {
		if v == GROUP_SIZE {
			commonChar = k
			break
		}
	}
	return commonChar
}

func part2(lines []string) rune {
	groups := make([][]string, len(lines)/GROUP_SIZE)
	for i := 0; i < len(lines); i += GROUP_SIZE {
		groups[i/GROUP_SIZE] = lines[i : i+GROUP_SIZE]
	}

	var sum rune = 0
	for _, group := range groups {
		commonChar := findCommonCharOfLines(group)
		sum += score(commonChar)
	}

	return sum
}

func main() {
	lines, _ := utils.ReadLines("input.txt")

	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}
