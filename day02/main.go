package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	entries := parseInput()
	var validCount int
	for _, e := range entries {
		var valid bool
		if string(e.pw[e.p.min-1]) == e.p.l {
			valid = !valid
		}
		if string(e.pw[e.p.max-1]) == e.p.l {
			valid = !valid
		}
		if valid {
			validCount++
		}
	}
	fmt.Println(validCount)
}

func main_1() {
	entries := parseInput()
	var valid int
outerloop:
	for _, e := range entries {
		var counter int
		for _, r := range e.pw {
			if string(r) == e.p.l {
				counter++
			}
			if counter > e.p.max {
				continue outerloop
			}
		}
		if counter < e.p.min {
			continue
		}
		valid++
	}
	fmt.Println(valid)
}

type entry struct {
	p  policy
	pw string
}

type policy struct {
	l   string
	min int
	max int
}

func parseInput() []entry {
	var entries []entry
	r := strings.NewReader(input)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		s := scanner.Text()
		ss := strings.Split(s, " ")
		policyRange := strings.Split(ss[0], "-")
		min, _ := strconv.Atoi(policyRange[0])
		max, _ := strconv.Atoi(policyRange[1])
		pol := policy{
			min: min,
			max: max,
			l:   strings.Trim(ss[1], ":"),
		}
		entry := entry{
			p:  pol,
			pw: ss[2],
		}
		entries = append(entries, entry)
	}
	return entries
}
