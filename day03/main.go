package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	input := parseInput()

	product := getTrees(1, 1, input)
	product *= getTrees(3, 1, input)
	product *= getTrees(5, 1, input)
	product *= getTrees(7, 1, input)
	product *= getTrees(1, 2, input)

	fmt.Println(product)
}

func getTrees(vX, vY int, input [][]string) int {
	limX := len(input[0])
	limY := len(input)

	var trees int
	for x, y := vX, vY; y < limY; y += vY {
		if input[y][x] == "#" {
			trees++
		}
		x = (x + vX) % limX
	}

	return trees
}

func main_1() {
	input := parseInput()

	vY := 1
	vX := 3
	limX := len(input[0])
	limY := len(input)

	var trees int
	for x, y := vX, vY; y < limY; y += vY {
		if input[y][x] == "#" {
			trees++
		}
		x = (x + vX) % limX
	}

	fmt.Println(trees)
}

func parseInput() [][]string {
	var grid [][]string
	r := strings.NewReader(input)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		ss := scanner.Text()
		var row []string
		for _, s := range ss {
			row = append(row, string(s))
		}
		grid = append(grid, row)
	}
	return grid
}
