package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	starting := strings.Split("17,1,3,16,19,0", ",")
	var spoken []int
	lastSeen := map[int]int{}  // number: index
	lastSeen2 := map[int]int{} // number: index
	for index, s := range starting {
		i, _ := strconv.Atoi(s)
		lastSeen[i] = index + 1
		spoken = append(spoken, i)
	}
	turn := len(lastSeen) + 1

	for turn <= 30000000 {
		last := spoken[len(spoken)-1]
		if _, ok := lastSeen2[last]; !ok {
			spoken = append(spoken, 0)
			if ls, ok := lastSeen[0]; ok {
				lastSeen2[0] = ls
			}
			lastSeen[0] = turn
		} else {
			age := lastSeen[last] - lastSeen2[last]
			spoken = append(spoken, age)
			if ls, ok := lastSeen[age]; ok {
				lastSeen2[age] = ls
			}
			lastSeen[age] = turn
		}
		turn++
	}
	fmt.Println(spoken[30000000-1])
}
