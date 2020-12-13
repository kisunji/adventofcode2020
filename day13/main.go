package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	ss := strings.Split(input, "\n")
	buses := strings.Split(ss[1], ",")

	var earliest int
	t := 0
outer:
	for {
		// tracks the product of all seen busIds
		increment := 1
		for i := 0; i < len(buses); i++ {
			if buses[i] == "x" {
				continue
			}
			busId, _ := strconv.Atoi(buses[i])
			if (t+i)%busId != 0 {
				t += increment
				continue outer
			}
			fmt.Printf("index %d time %d\n", i, t)
			increment *= busId
		}
		earliest = t
		break outer
	}
	fmt.Printf("earliest: %d", earliest)
}

func main_1() {
	ss := strings.Split(input, "\n")
	time, _ := strconv.Atoi(ss[0])
	buses := strings.Split(ss[1], ",")
	var fastestTime int
	var earliestBusId int
	for _, bus := range buses {
		if bus == "x" {
			continue
		}
		busId, _ := strconv.Atoi(bus)
		i := time / busId
		if i*busId == time {
			fastestTime = time
			earliestBusId = busId
			break
		}
		i++
		if i*busId < fastestTime || fastestTime == 0 {
			fastestTime = i * busId
			earliestBusId = busId
		}
	}
	fmt.Println((fastestTime - time) * earliestBusId)
}

const sample = `939
7,13,x,x,59,x,31,19`

const input = `1000303
41,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,37,x,x,x,x,x,541,x,x,x,x,x,x,x,23,x,x,x,x,13,x,x,x,17,x,x,x,x,x,x,x,x,x,x,x,29,x,983,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,19`
